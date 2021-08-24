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
          <th>操作系统</th>
          <th>在线时间</th>
          <th>负载占用率</th>
          <th>CPU 占用率</th>
          <th>内存占用率</th>
          <th>SWAP 占用率</th>
          <th>磁盘占用率</th>
          <th></th>
        </template>

        <template v-slot:default="row">
          <th scope="row">
            <div class="media align-items-center">
              <a href="#" class="avatar rounded-circle mr-3">
                <img alt="Image placeholder" :src="row.item.countryImg" />
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
            {{ row.item.os }}
          </td>
          <td>
            {{ row.item.uptime }}
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
export default {
  name: "nodes-table",
  props: {
    type: {
      type: String,
    },
    title: String,
  },
  data() {
    return {
      tableData: [
        {
          countryImg: "img/theme/bootstrap.jpg",
          friendlyName: "Server",
          statusType: "warning",
          status: "离线",
          os: "Unknown",
          uptime: "0",
          load: 0,
          loadStatusType: "failure",
          cpu: 0,
          cpuStatusType: "failure",
          mem: 0,
          memStatusType: "failure",
          swap: 0,
          swapStatusType: "failure",
          disk: 0,
          diskStatusType: "failure",
        },
      ],
    };
  },
};
</script>
<style></style>
