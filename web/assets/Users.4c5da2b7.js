import"./request.a57216d8.js";import{g as e,s as a}from"./user.36613f22.js";import{s as t,u as l,g as r,a as s}from"./cluster.e0a8b9ec.js";import{d as n,C as u,B as d,r as o,o as c,e as i,f as p,g as f,w as m,c as h,h as g,L as b,D as _,p as y,j as C,i as k}from"./vendor.da3fc754.js";import{_ as v}from"./index.939ea1aa.js";const w={name:"users",setup(){n();const o=u({username:"",pageIndex:1,pageSize:10}),c=d([]),i=d(0),p=()=>{e(o).then((e=>{c.value=e.data.users,i.value=e.data.count||50}))};p();const f=d(!1);let m=u({name:"",raft_id:""}),h=0;const g=d([]),y=e=>{r().then((e=>{0===e.status?s(o).then((a=>{if(0===a.status){for(let t=0;t<e.data.count;t++){let l=!1;for(let r=0;r<a.data.count;r++)if(e.data.cluster[t].raft_id===a.data.cluster[r].raft_id){l=!0,e.data.cluster[t].auth=1;break}l||(e.data.cluster[t].auth=0)}g.value=e.data.cluster}else _.error("获取失败")})):_.error("获取失败")}))};return{query:o,tableData:c,pageTotal:i,handleSearch:()=>{o.pageIndex=1,p()},handlePageChange:e=>{o.pageIndex=e,p()},clusterVisible:f,form:m,clusters:g,manageCluster:e=>{y(),h=e,f.value=!0},setAdmin:e=>{b.confirm("确定要设为管理员吗？","提示",{type:"warning"}).then((()=>{a({id:e}).then((e=>{0===e.status?_.success("操作成功"):_.error("操作失败")})),p()})).catch((()=>{}))},changeClusterAuth:(e,a)=>{b.confirm("确定要更改权限吗？","提示",{type:"warning"}).then((()=>{0===a?t({user_id:h,raft_id:e}).then((e=>{0===e.status?_.success("操作成功"):_.error("操作失败")})):l({user_id:h,raft_id:e}).then((e=>{0===e.status?_.success("操作成功"):_.error("操作失败")})),y()})).catch((()=>{}))}}}},x={class:"crumbs"},V=(e=>(y("data-v-42567fcb"),e=e(),C(),e))((()=>p("i",{class:"el-icon-lx-cascades"},null,-1))),I=k(" 用户管理 "),j={class:"container"},q={class:"handle-box"},D=k("搜索"),z={key:0},A={key:1},S=k("管理集群权限"),T=k("设为管理员"),P={class:"pagination"},U={key:0},B={key:1},L=k("改变集群权限"),E={class:"dialog-footer"},F=k("确 定");var G=v(w,[["render",function(e,a,t,l,r,s){const n=o("el-breadcrumb-item"),u=o("el-breadcrumb"),d=o("el-input"),b=o("el-button"),_=o("el-table-column"),y=o("el-tag"),C=o("el-table"),k=o("el-pagination"),v=o("el-dialog");return c(),i("div",null,[p("div",x,[f(u,{separator:"/"},{default:m((()=>[f(n,null,{default:m((()=>[V,I])),_:1})])),_:1})]),p("div",j,[p("div",q,[f(d,{modelValue:l.query.username,"onUpdate:modelValue":a[0]||(a[0]=e=>l.query.username=e),placeholder:"用户名",class:"handle-input mr10"},null,8,["modelValue"]),f(b,{type:"primary",icon:"el-icon-search",onClick:l.handleSearch},{default:m((()=>[D])),_:1},8,["onClick"])]),f(C,{data:l.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:m((()=>[f(_,{prop:"id",label:"ID",width:"55",align:"center"}),f(_,{prop:"username",label:"用户名"}),f(_,{prop:"is_admin",label:"身份"},{default:m((e=>[f(y,{effect:"dark",type:1===e.row.is_admin?"danger":"success"},{default:m((()=>[1===e.row.is_admin?(c(),i("p",z," 管理员 ")):(c(),i("p",A," 普通用户 "))])),_:2},1032,["type"])])),_:1}),f(_,{label:"操作",align:"center"},{default:m((e=>[f(b,{type:"text",icon:"el-icon-s-grid",onClick:a=>l.manageCluster(e.row.id)},{default:m((()=>[S])),_:2},1032,["onClick"]),0===e.row.is_admin?(c(),h(b,{key:0,type:"text",icon:"el-icon-s-custom",onClick:a=>l.setAdmin(e.row.id)},{default:m((()=>[T])),_:2},1032,["onClick"])):g("",!0)])),_:1})])),_:1},8,["data"]),p("div",P,[f(k,{background:"",layout:"total, prev, pager, next","current-page":l.query.pageIndex,"page-size":l.query.pageSize,total:l.pageTotal,onCurrentChange:l.handlePageChange},null,8,["current-page","page-size","total","onCurrentChange"])])]),f(v,{title:"用户集群权限管理",modelValue:l.clusterVisible,"onUpdate:modelValue":a[2]||(a[2]=e=>l.clusterVisible=e),width:"80%"},{footer:m((()=>[p("span",E,[f(b,{type:"primary",onClick:a[1]||(a[1]=e=>l.clusterVisible=!1)},{default:m((()=>[F])),_:1})])])),default:m((()=>[f(C,{data:l.clusters,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:m((()=>[f(_,{prop:"raft_id",label:"集群ID"}),f(_,{prop:"address",label:"地址"}),f(_,{prop:"auth",label:"权限"},{default:m((e=>[f(y,{effect:"dark",type:1===e.row.auth?"success":"info"},{default:m((()=>[1===e.row.auth?(c(),i("p",U," 有 ")):(c(),i("p",B," 无 "))])),_:2},1032,["type"])])),_:1}),f(_,{label:"操作",align:"center"},{default:m((e=>[f(b,{type:"text",icon:"el-icon-s-grid",onClick:a=>l.changeClusterAuth(e.row.raft_id,e.row.auth)},{default:m((()=>[L])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])])),_:1},8,["modelValue"])])}],["__scopeId","data-v-42567fcb"]]);export{G as default};
