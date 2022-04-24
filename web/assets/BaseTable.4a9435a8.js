import{A as e,z as a,p as l,a as t,L as d,B as n,r as s,o,c as r,f as u,h as c,t as i,w as p}from"./vendor.72ef4fca.js";import{s as m}from"./request.111ca7cf.js";const b={name:"basetable",setup(){const l=e({address:"",name:"",pageIndex:1,pageSize:10}),t=a([]),s=a(0),o=()=>{(e=>m({url:"./table.json",method:"get",params:e}))(l).then((e=>{t.value=e.list,s.value=e.pageTotal||50}))};o();const r=a(!1);let u=e({name:"",address:""}),c=-1;return{query:l,tableData:t,pageTotal:s,editVisible:r,form:u,handleSearch:()=>{l.pageIndex=1,o()},handlePageChange:e=>{l.pageIndex=e,o()},handleDelete:e=>{d.confirm("确定要删除吗？","提示",{type:"warning"}).then((()=>{n.success("删除成功"),t.value.splice(e,1)})).catch((()=>{}))},handleEdit:(e,a)=>{c=e,Object.keys(u).forEach((e=>{u[e]=a[e]})),r.value=!0},saveEdit:()=>{r.value=!1,n.success(`修改第 ${c+1} 行成功`),Object.keys(u).forEach((e=>{t.value[c][e]=u[e]}))}}}},f=p();l("data-v-5505848c");const h={class:"crumbs"},g=u("i",{class:"el-icon-lx-cascades"},null,-1),y=c(" 基础表格 "),v={class:"container"},_={class:"handle-box"},V=c("搜索"),w=c("编辑 "),C=c("删除"),k={class:"pagination"},x={class:"dialog-footer"},q=c("取 消"),E=c("确 定");t();const I=f(((e,a,l,t,d,n)=>{const p=s("el-breadcrumb-item"),m=s("el-breadcrumb"),b=s("el-option"),I=s("el-select"),j=s("el-input"),z=s("el-button"),D=s("el-table-column"),U=s("el-image"),S=s("el-tag"),T=s("el-table"),$=s("el-pagination"),O=s("el-form-item"),P=s("el-form"),A=s("el-dialog");return o(),r("div",null,[u("div",h,[u(m,{separator:"/"},{default:f((()=>[u(p,null,{default:f((()=>[g,y])),_:1})])),_:1})]),u("div",v,[u("div",_,[u(I,{modelValue:t.query.address,"onUpdate:modelValue":a[1]||(a[1]=e=>t.query.address=e),placeholder:"地址",class:"handle-select mr10"},{default:f((()=>[u(b,{key:"1",label:"广东省",value:"广东省"}),u(b,{key:"2",label:"湖南省",value:"湖南省"})])),_:1},8,["modelValue"]),u(j,{modelValue:t.query.name,"onUpdate:modelValue":a[2]||(a[2]=e=>t.query.name=e),placeholder:"用户名",class:"handle-input mr10"},null,8,["modelValue"]),u(z,{type:"primary",icon:"el-icon-search",onClick:t.handleSearch},{default:f((()=>[V])),_:1},8,["onClick"])]),u(T,{data:t.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:f((()=>[u(D,{prop:"id",label:"ID",width:"55",align:"center"}),u(D,{prop:"name",label:"用户名"}),u(D,{label:"账户余额"},{default:f((e=>[c("￥"+i(e.row.money),1)])),_:1}),u(D,{label:"头像(查看大图)",align:"center"},{default:f((e=>[u(U,{class:"table-td-thumb",src:e.row.thumb,"preview-src-list":[e.row.thumb]},null,8,["src","preview-src-list"])])),_:1}),u(D,{prop:"address",label:"地址"}),u(D,{label:"状态",align:"center"},{default:f((e=>[u(S,{type:"成功"===e.row.state?"success":"失败"===e.row.state?"danger":""},{default:f((()=>[c(i(e.row.state),1)])),_:2},1032,["type"])])),_:1}),u(D,{prop:"date",label:"注册时间"}),u(D,{label:"操作",width:"180",align:"center"},{default:f((e=>[u(z,{type:"text",icon:"el-icon-edit",onClick:a=>t.handleEdit(e.$index,e.row)},{default:f((()=>[w])),_:2},1032,["onClick"]),u(z,{type:"text",icon:"el-icon-delete",class:"red",onClick:a=>t.handleDelete(e.$index,e.row)},{default:f((()=>[C])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),u("div",k,[u($,{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal,onCurrentChange:t.handlePageChange},null,8,["current-page","page-size","total","onCurrentChange"])])]),u(A,{title:"编辑",modelValue:t.editVisible,"onUpdate:modelValue":a[6]||(a[6]=e=>t.editVisible=e),width:"30%"},{footer:f((()=>[u("span",x,[u(z,{onClick:a[5]||(a[5]=e=>t.editVisible=!1)},{default:f((()=>[q])),_:1}),u(z,{type:"primary",onClick:t.saveEdit},{default:f((()=>[E])),_:1},8,["onClick"])])])),default:f((()=>[u(P,{"label-width":"70px"},{default:f((()=>[u(O,{label:"用户名"},{default:f((()=>[u(j,{modelValue:t.form.name,"onUpdate:modelValue":a[3]||(a[3]=e=>t.form.name=e)},null,8,["modelValue"])])),_:1}),u(O,{label:"地址"},{default:f((()=>[u(j,{modelValue:t.form.address,"onUpdate:modelValue":a[4]||(a[4]=e=>t.form.address=e)},null,8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"])])}));b.render=I,b.__scopeId="data-v-5505848c";export default b;
