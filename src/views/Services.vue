<template>
  <div>
    <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
    </base-header>

    <div class="container-fluid mt--7">
      <div class="row">
        <div class="col">
          <nodes-table title="节点监控"></nodes-table>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import NodesTable from "./Tables/NodesTable";
export default {
  name: "tables",
  components: {
    NodesTable,
  },
};
</script>
<style></style>
