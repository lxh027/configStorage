import{g as e,a}from"./index.d0a782d5.js";import{s as t,u as l,g as r,a as s}from"./cluster.0aebc448.js";import{f as n,B as d,A as u,p as o,b as c,L as i,C as p,r as f,o as m,c as h,g,m as b,w as _,h as y}from"./vendor.971a53c8.js";const C={name:"users",setup(){n();const o=d({username:"",pageIndex:1,pageSize:10}),c=u([]),f=u(0),m=()=>{e(o).then((e=>{c.value=e.data.users,f.value=e.data.count||50}))};m();const h=u(!1);let g=d({name:"",raft_id:""}),b=0;const _=u([]),y=e=>{r().then((a=>{0===a.status?s({user_id:e}).then((e=>{if(0===e.status){for(let t=0;t<a.data.count;t++){let l=!1;for(let r=0;r<e.data.count;r++)if(a.data.cluster[t].raft_id===e.data.cluster[r].raft_id){l=!0,a.data.cluster[t].auth=1;break}l||(a.data.cluster[t].auth=0)}_.value=a.data.cluster}else p.error("获取失败")})):p.error("获取失败")}))};return{query:o,tableData:c,pageTotal:f,handleSearch:()=>{o.pageIndex=1,m()},handlePageChange:e=>{o.pageIndex=e,m()},clusterVisible:h,form:g,clusters:_,manageCluster:e=>{y(e),b=e,h.value=!0},setAdmin:e=>{i.confirm("确定要设为管理员吗？","提示",{type:"warning"}).then((()=>{a({id:e}).then((e=>{0===e.status?p.success("操作成功"):p.error("操作失败")})),m()})).catch((()=>{}))},changeClusterAuth:(e,a)=>{i.confirm("确定要更改权限吗？","提示",{type:"warning"}).then((()=>{0===a?t({user_id:b,raft_id:e}).then((e=>{0===e.status?p.success("操作成功"):p.error("操作失败")})):l({user_id:b,raft_id:e}).then((e=>{0===e.status?p.success("操作成功"):p.error("操作失败")})),y(b)})).catch((()=>{}))}}}},k=_();o("data-v-f5040f2a");const w={class:"crumbs"},v=g("i",{class:"el-icon-lx-cascades"},null,-1),x=y(" 用户管理 "),V={class:"container"},I={class:"handle-box"},q=y("搜索"),A={key:0},z={key:1},D=y("管理集群权限"),S=y("设为管理员"),T={class:"pagination"},j={key:0},P={key:1},U=y("改变集群权限"),B={class:"dialog-footer"},L=y("确 定");c();const E=k(((e,a,t,l,r,s)=>{const n=f("el-breadcrumb-item"),d=f("el-breadcrumb"),u=f("el-input"),o=f("el-button"),c=f("el-table-column"),i=f("el-tag"),p=f("el-table"),_=f("el-pagination"),y=f("el-dialog");return m(),h("div",null,[g("div",w,[g(d,{separator:"/"},{default:k((()=>[g(n,null,{default:k((()=>[v,x])),_:1})])),_:1})]),g("div",V,[g("div",I,[g(u,{modelValue:l.query.username,"onUpdate:modelValue":a[1]||(a[1]=e=>l.query.username=e),placeholder:"用户名",class:"handle-input mr10"},null,8,["modelValue"]),g(o,{type:"primary",icon:"el-icon-search",onClick:l.handleSearch},{default:k((()=>[q])),_:1},8,["onClick"])]),g(p,{data:l.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:k((()=>[g(c,{prop:"id",label:"ID",width:"55",align:"center"}),g(c,{prop:"username",label:"用户名"}),g(c,{prop:"is_admin",label:"身份"},{default:k((e=>[g(i,{effect:"dark",type:1===e.row.is_admin?"danger":"success"},{default:k((()=>[1===e.row.is_admin?(m(),h("p",A," 管理员 ")):(m(),h("p",z," 普通用户 "))])),_:2},1032,["type"])])),_:1}),g(c,{label:"操作",align:"center"},{default:k((e=>[g(o,{type:"text",icon:"el-icon-s-grid",onClick:a=>l.manageCluster(e.row.id)},{default:k((()=>[D])),_:2},1032,["onClick"]),0===e.row.is_admin?(m(),h(o,{key:0,type:"text",icon:"el-icon-s-custom",onClick:a=>l.setAdmin(e.row.id)},{default:k((()=>[S])),_:2},1032,["onClick"])):b("",!0)])),_:1})])),_:1},8,["data"]),g("div",T,[g(_,{background:"",layout:"total, prev, pager, next","current-page":l.query.pageIndex,"page-size":l.query.pageSize,total:l.pageTotal,onCurrentChange:l.handlePageChange},null,8,["current-page","page-size","total","onCurrentChange"])])]),g(y,{title:"用户集群权限管理",modelValue:l.clusterVisible,"onUpdate:modelValue":a[3]||(a[3]=e=>l.clusterVisible=e),width:"80%"},{footer:k((()=>[g("span",B,[g(o,{type:"primary",onClick:a[2]||(a[2]=e=>l.clusterVisible=!1)},{default:k((()=>[L])),_:1})])])),default:k((()=>[g(p,{data:l.clusters,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:k((()=>[g(c,{prop:"raft_id",label:"集群ID"}),g(c,{prop:"address",label:"地址"}),g(c,{prop:"auth",label:"权限"},{default:k((e=>[g(i,{effect:"dark",type:1===e.row.auth?"success":"info"},{default:k((()=>[1===e.row.auth?(m(),h("p",j," 有 ")):(m(),h("p",P," 无 "))])),_:2},1032,["type"])])),_:1}),g(c,{label:"操作",align:"center"},{default:k((e=>[g(o,{type:"text",icon:"el-icon-s-grid",onClick:a=>l.changeClusterAuth(e.row.raft_id,e.row.auth)},{default:k((()=>[U])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])])),_:1},8,["modelValue"])])}));C.render=E,C.__scopeId="data-v-f5040f2a";export default C;
