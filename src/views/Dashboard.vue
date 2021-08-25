<template>
  <div>
    <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
      <div class="row">
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="正常节点"
            type="gradient-green"
            :sub-title="onlineNodesCnt"
            icon="ni ni-button-play"
            class="mb-4 mb-xl-0"
          >
            <template v-slot:footer>
              <span class="text-success mr-2">
                <i class="fa fa-circle"></i>
              </span>
              <span class="text-nowrap">近5分钟在线</span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="异常节点"
            type="gradient-red"
            :sub-title="offlineNodesCnt"
            icon="ni ni-button-pause"
            class="mb-4 mb-xl-0"
          >
            <template v-slot:footer>
              <span class="text-warning mr-2">
                <i class="fa fa-circle"></i>
              </span>
              <span class="text-nowrap">近5分钟离线</span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="正常服务"
            type="gradient-green"
            :sub-title="onlineServicesCnt"
            icon="ni ni-button-play"
            class="mb-4 mb-xl-0"
          >
            <template v-slot:footer>
              <span class="text-success mr-2">
                <i class="fa fa-circle"></i>
              </span>
              <span class="text-nowrap">近5分钟在线</span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="异常服务"
            type="gradient-red"
            :sub-title="offlineServicesCnt"
            icon="ni ni-button-pause"
            class="mb-4 mb-xl-0"
          >
            <template v-slot:footer>
              <span class="text-warning mr-2">
                <i class="fa fa-circle"></i>
              </span>
              <span class="text-nowrap">近5分钟离线</span>
            </template>
          </stats-card>
        </div>
      </div>
    </base-header>

    <div class="container-fluid mt--7">
      <!--Charts-->
      <div class="row">
        <div class="col-xl-8 mb-5 mb-xl-0">
          <card type="default" header-classes="bg-transparent">
            <template v-slot:header>
              <div class="row align-items-center">
                <div class="col">
                  <h6 class="text-light text-uppercase ls-1 mb-1">负载概览</h6>
                  <h5 class="h3 text-white mb-0">5分钟负载（%）</h5>
                </div>
                <div class="col">
                  <ul class="nav nav-pills justify-content-end">
                    <li class="nav-item mr-2 mr-md-0">
                      <a
                        class="nav-link py-2 px-3"
                        href="#"
                        :class="{ active: nodesLoadChart.activeIndex === 0 }"
                        @click.prevent="initNodesLoadChart(0)"
                      >
                        <span class="d-none d-md-block">周视图</span>
                        <span class="d-md-none">周</span>
                      </a>
                    </li>
                    <li class="nav-item">
                      <a
                        class="nav-link py-2 px-3"
                        href="#"
                        :class="{ active: nodesLoadChart.activeIndex === 1 }"
                        @click.prevent="initNodesLoadChart(1)"
                      >
                        <span class="d-none d-md-block">日视图</span>
                        <span class="d-md-none">日</span>
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
            </template>
            <div class="chart-area">
              <canvas :height="350" :id="nodesLoadChartID"></canvas>
            </div>
          </card>
        </div>

        <div class="col-xl-4">
          <card header-classes="bg-transparent">
            <template v-slot:header>
              <div class="row align-items-center">
                <div class="col">
                  <h6 class="text-uppercase text-muted ls-1 mb-1">
                    Performance
                  </h6>
                  <h5 class="h3 mb-0">Total orders</h5>
                </div>
              </div>
            </template>
            <div class="chart-area">
              <canvas :height="350" :id="ordersChartID"></canvas>
            </div>
          </card>
        </div>
      </div>
      <!--End charts-->

      <!--Tables-->
      <div class="row mt-5">
        <div class="col-xl-8 mb-5 mb-xl-0">
          <page-visits-table></page-visits-table>
        </div>
        <div class="col-xl-4">
          <social-traffic-table></social-traffic-table>
        </div>
      </div>
      <!--End tables-->
    </div>
  </div>
</template>
<script>
// Charts
import { ordersChart } from "@/components/Charts/Chart";
import Chart from "chart.js";
import axios from "axios";

import PageVisitsTable from "./Dashboard/PageVisitsTable";
import SocialTrafficTable from "./Dashboard/SocialTrafficTable";
let chart;

export default {
  components: {
    PageVisitsTable,
    SocialTrafficTable,
  },
  data() {
    return {
      onlineNodesCnt: null,
      offlineNodesCnt: null,
      onlineServicesCnt: "0",
      offlineServicesCnt: "0",
      nodesLoadChartID: "nodesLoadChart",
      ordersChartID: "ordersChart",
      nodesLoadChart: {
        allData: [
          [0, 20, 10, 30, 15, 40, 20, 60, 60],
          [0, 20, 5, 25, 10, 30, 15, 40, 40],
        ],
        activeIndex: 0,
      },
    };
  },
  methods: {
    getNodesStat() {
      axios
        .get(process.env.VUE_APP_FIREBASE_RTDB_URL + "/nodes.json?shallow=true")
        .then(async (totalNodes) => {
          await axios
            .get(
              `${
                process.env.VUE_APP_FIREBASE_RTDB_URL
              }/nodes.json?orderBy="LastUpdate"&startAt="${(
                Math.round(Date.now() / 1000) - 305
              ).toString()}"&endAt="${(
                Math.round(Date.now() / 1000) + 5
              ).toString()}"`
            )
            .then((onlineNodes) => {
              this.onlineNodesCnt = Object.keys(
                onlineNodes.data
              ).length.toString();
            });
          this.offlineNodesCnt = (
            Object.keys(totalNodes.data).length - parseInt(this.onlineNodesCnt)
          ).toString();
        });
    },
    initNodesLoadChart(index) {
      chart.destroy();
      chart = new Chart(
        document.getElementById(this.nodesLoadChartID).getContext("2d"),
        {
          type: "line",
          data: {
            labels: ["May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
            datasets: [
              {
                label: "Performance",
                tension: 0.4,
                borderWidth: 4,
                borderColor: "#5e72e4",
                pointRadius: 0,
                backgroundColor: "transparent",
                data: this.nodesLoadChart.allData[index],
              },
            ],
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            legend: {
              display: false,
            },
            tooltips: {
              enabled: true,
              mode: "index",
              intersect: false,
            },
            scales: {
              yAxes: [
                {
                  barPercentage: 1.6,
                  gridLines: {
                    drawBorder: false,
                    color: "rgba(29,140,248,0.0)",
                    zeroLineColor: "transparent",
                  },
                  ticks: {
                    padding: 0,
                    fontColor: "#8898aa",
                    fontSize: 13,
                    fontFamily: "Open Sans",
                  },
                },
              ],
              xAxes: [
                {
                  barPercentage: 1.6,
                  gridLines: {
                    drawBorder: false,
                    color: "rgba(29,140,248,0.0)",
                    zeroLineColor: "transparent",
                  },
                  ticks: {
                    padding: 10,
                    fontColor: "#8898aa",
                    fontSize: 13,
                    fontFamily: "Open Sans",
                  },
                },
              ],
            },
            layout: {
              padding: 0,
            },
          },
        }
      );
      this.nodesLoadChart.activeIndex = index;
    },
  },
  mounted() {
    this.getNodesStat();
    chart = new Chart(
      document.getElementById(this.nodesLoadChartID).getContext("2d"),
      {
        type: "line",
        data: {
          labels: ["May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
          datasets: [
            {
              label: "Performance",
              tension: 0.4,
              borderWidth: 4,
              borderColor: "#5e72e4",
              pointRadius: 0,
              backgroundColor: "transparent",
              data: this.nodesLoadChart.allData[1],
            },
          ],
        },
        options: {
          responsive: true,
          maintainAspectRatio: false,
          legend: {
            display: false,
          },
          tooltips: {
            enabled: true,
            mode: "index",
            intersect: false,
          },
          scales: {
            yAxes: [
              {
                barPercentage: 1.6,
                gridLines: {
                  drawBorder: false,
                  color: "rgba(29,140,248,0.0)",
                  zeroLineColor: "transparent",
                },
                ticks: {
                  padding: 0,
                  fontColor: "#8898aa",
                  fontSize: 13,
                  fontFamily: "Open Sans",
                },
              },
            ],
            xAxes: [
              {
                barPercentage: 1.6,
                gridLines: {
                  drawBorder: false,
                  color: "rgba(29,140,248,0.0)",
                  zeroLineColor: "transparent",
                },
                ticks: {
                  padding: 10,
                  fontColor: "#8898aa",
                  fontSize: 13,
                  fontFamily: "Open Sans",
                },
              },
            ],
          },
          layout: {
            padding: 0,
          },
        },
      }
    );
    ordersChart.createChart(this.ordersChartID);
  },
};
</script>
<style></style>
