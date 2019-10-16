<template>
  <div class="duty">
      <el-table
      :data="tableData"
      style="width: 100%"
      :default-sort = "{prop: 'date', order: 'descending'}"
      >
      <el-table-column
        prop="Id"
        label="Id"
        width="180">
      </el-table-column>
      <el-table-column
        prop="Begindate"
        label="Begindate"
        sortable
        width="240">
      </el-table-column>
            <el-table-column
        prop="Enddate"
        label="Enddate"
        sortable
        width="240">
      </el-table-column>
                  <el-table-column
        prop="Request_person"
        label="Request_person"
        width="180">
      </el-table-column>
                  <el-table-column
        prop="Alert_person"
        label="Alert_person"
        width="180">
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
// @ is an alias to /src
import * as API from '@/api/dbinstance/';


export default {
  name: 'duty',
  data() {
return {
 tableData :[],
};
},
methods :{
  load(){
    API.getDutys().then((res) => {
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
}
</script>