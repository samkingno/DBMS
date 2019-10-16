<template>
  <div class="dbinstances">
      <el-table
      :data="tableData"
      style="width: 100%"
      :default-sort = "{prop: 'date', order: 'descending'}"
      >
       <el-table-column
         prop="id"
         label="ID"
         sortable
         width="180">
       </el-table-column>
      <el-table-column
        prop="db_host_id"
        label="db_host_id"
        sortable
        width="180">
      </el-table-column>
      <el-table-column
        prop="hostname"
        label="hostname"
        sortable
        width="280">
      </el-table-column>
           <el-table-column
        prop="port"
        label="port"
        sortable
        width="180">
      </el-table-column>
           <el-table-column
        prop="status"
        label="status"
        sortable
        width="180">
      </el-table-column>
               <el-table-column
        prop="cluster_product_name"
        label="cluster_product_name"
        sortable
        width="200">
      </el-table-column>
                <el-table-column
        prop="cluster_product_owner"
        label="cluster_product_owner"
        sortable
        width="200">
      </el-table-column>
                <el-table-column
        prop="pg_version"
        label="pg_version"
        sortable
        width="180">
      </el-table-column>
                <el-table-column
        prop="master_slave"
        label="master_slave"
        :formatter="formatter"
        width="180">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
// @ is an alias to /src
import * as API from '@/api/dbinstance/';


export default {
  name: 'dbinstances',
data() {
return {
 tableData :[],
};
},
methods :{
  load(){
    API.getDbinstances().then((res) => {
      this.tableData = res.data;
    });
  },
        formatter(row) {
        return row.master_slave;
      }
},
components:{
  },
  beforeMount() {
    this.load();
  },
};
</script>
