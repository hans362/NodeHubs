<template>
  <div class="card shadow" :class="type === 'dark' ? 'bg-default' : ''">
    <div
      class="card-header border-0"
      :class="type === 'dark' ? 'bg-transparent' : ''"
    >
      <div class="row align-items-center">
        <div class="col">
          <h3 class="mb-0" :class="type === 'dark' ? 'text-white' : ''">
            {{ title }}
          </h3>
        </div>
      </div>
    </div>

    <div class="table-responsive">
      <base-table
        class="table align-items-center table-flush"
        :class="type === 'dark' ? 'table-dark' : ''"
        :thead-classes="type === 'dark' ? 'thead-dark' : 'thead-light'"
        tbody-classes="list"
        :data="tableData"
      >
        <template v-slot:columns>
          <th>节点名称</th>
          <th>基本信息</th>
          <th>负载占用</th>
          <th>CPU 占用率</th>
          <th>内存占用率</th>
          <th>SWAP 占用率</th>
          <th>磁盘占用率</th>
        </template>

        <template v-slot:default="row">
          <th scope="row">
            <div class="media align-items-center">
              <a href="#" class="mr-3">
                <country-flag :country="row.item.countryFlag" size="big" />
              </a>
              <div class="media-body">
                <span class="name mb-0 text-sm">{{
                  row.item.friendlyName
                }}</span>
                <br />
                <badge class="badge-dot mr-4" :type="row.item.statusType">
                  <i :class="`bg-${row.item.statusType}`"></i>
                  <span class="status">{{ row.item.status }}</span>
                </badge>
              </div>
            </div>
          </th>
          <td>
            {{ row.item.cpuModel }}
            <br />
            {{ row.item.os }}
            <br />
            {{ row.item.ip }}
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.item.load }}%</span>
              <div>
                <base-progress
                  :type="row.item.loadStatusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.item.load"
                />
              </div>
            </div>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.item.cpu }}%</span>
              <div>
                <base-progress
                  :type="row.item.cpuStatusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.item.cpu"
                />
              </div>
            </div>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.item.mem }}%</span>
              <div>
                <base-progress
                  :type="row.item.memStatusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.item.mem"
                />
              </div>
            </div>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.item.swap }}%</span>
              <div>
                <base-progress
                  :type="row.item.swapStatusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.item.swap"
                />
              </div>
            </div>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.item.disk }}%</span>
              <div>
                <base-progress
                  :type="row.item.diskStatusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.item.disk"
                />
              </div>
            </div>
          </td>
        </template>
      </base-table>
    </div>
  </div>
</template>
<script>
import countryFlag from "vue-country-flag-next";
import axios from "axios";

export default {
  name: "nodes-table",
  components: {
    countryFlag,
  },
  props: {
    type: {
      type: String,
    },
    title: String,
  },
  data() {
    return {
      tableData: [],
    };
  },
  methods: {
    getStat(lastUpdate) {
      if (lastUpdate <= Math.round(Date.now() / 1000) - 305) {
        return false;
      } else {
        return true;
      }
    },
    getStatusType(percent) {
      if (percent <= 60) {
        return "success";
      } else if (percent <= 80) {
        return "warning";
      } else {
        return "danger";
      }
    },
    getNodesList() {
      axios
        .get(`${process.env.VUE_APP_FIREBASE_RTDB_URL}/nodes.json?shallow=true`)
        .then((nodesInfo) => {
          Object.keys(nodesInfo.data).forEach((nodeId) => {
            axios
              .get(
                `${process.env.VUE_APP_FIREBASE_RTDB_URL}/nodes/${nodeId}.json?orderBy="Timestamp"&limitToLast=1`
              )
              .then((rawInfo) => {
                var nodeInfo = rawInfo.data[Object.keys(rawInfo.data)[0]];
                console.log(nodeInfo);
                console.log(nodeInfo.Load.split(" ")[1] / nodeInfo.CPUCores);
                var nodeStat = this.getStat(parseInt(nodeInfo.Timestamp));
                var node = {
                  countryFlag: JSON.parse(
                    nodeInfo.IPInfo
                  ).countryCode.toLowerCase(),
                  friendlyName: nodeInfo.FriendlyName,
                  statusType: nodeStat ? "success" : "danger",
                  status: nodeStat ? "在线" : "离线",
                  os: `${
                    nodeInfo.Platform.charAt(0).toUpperCase() +
                    nodeInfo.Platform.slice(1)
                  } ${nodeInfo.PlatformVersion}`,
                  ip: JSON.parse(nodeInfo.IPInfo).query.toString(),
                  load: nodeStat
                    ? Math.round(
                        (nodeInfo.Load.split(" ")[1] / nodeInfo.CPUCores) * 100
                      )
                    : 0,
                  loadStatusType: this.getStatusType(
                    Math.round(
                      (nodeInfo.Load.split(" ")[1] / nodeInfo.CPUCores) * 100
                    )
                  ),
                  cpuModel: nodeInfo.CPUModel,
                  cpu: nodeStat ? parseInt(nodeInfo.CPUPercent) : 0,
                  cpuStatusType: this.getStatusType(
                    parseInt(nodeInfo.CPUPercent)
                  ),
                  mem: nodeStat ? parseInt(nodeInfo.MemPercent) : 0,
                  memStatusType: this.getStatusType(
                    parseInt(nodeInfo.MemPercent)
                  ),
                  swap: nodeStat ? parseInt(nodeInfo.SwapPercent) : 0,
                  swapStatusType: this.getStatusType(
                    parseInt(nodeInfo.SwapPercent)
                  ),
                  disk: nodeStat ? parseInt(nodeInfo.DiskPercent) : 0,
                  diskStatusType: this.getStatusType(
                    parseInt(nodeInfo.DiskPercent)
                  ),
                };
                this.tableData.push(node);
              });
          });
        });
    },
  },
  mounted() {
    this.getNodesList();
  },
};
</script>
<style></style>
