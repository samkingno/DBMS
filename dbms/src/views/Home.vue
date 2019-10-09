<template>
  <div class="home">
      <el-table
      :data="tableData"
      style="width: 100%">
       <el-table-column
         prop="id"
         label="ID"
         width="180">
       </el-table-column>
      <el-table-column
        prop="db_host_id"
        label="db_host_id"
        width="180">
      </el-table-column>
      <el-table-column
        prop="hostname"
        label="hostname"
        width="280">
      </el-table-column>
           <el-table-column
        prop="port"
        label="port"
        width="180">
      </el-table-column>
           <el-table-column
        prop="status"
        label="status"
        width="180">
      </el-table-column>
               <el-table-column
        prop="cluster_product_name"
        label="cluster_product_name"
        width="180">
      </el-table-column>
                <el-table-column
        prop="cluster_product_owner"
        label="cluster_product_owner"
        width="180">
      </el-table-column>
                <el-table-column
        prop="pg_version"
        label="pg_version"
        width="180">
      </el-table-column>
                <el-table-column
        prop="master_slave"
        label="master_slave"
        width="180">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
// @ is an alias to /src
import * as API from '@/api/dbinstance/';


export default {
  name: 'home',
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
},
components:{
  },
  beforeMount() {
    this.load();
  },
};
</script>
