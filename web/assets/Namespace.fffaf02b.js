import{s as e}from"./request.111ca7cf.js";import{a}from"./cluster.7b371fde.js";import{e as l,A as t,z as o,p as d,a as n,L as r,B as s,r as u,o as i,c as p,f as c,h as m,t as f,g as h,F as _,j as g,w as y}from"./vendor.72ef4fca.js";const b={name:"namespace",setup(){const d=l(),n=t({name:"",pageIndex:1,pageSize:10}),u=o([]),i=o(0),p=()=>{(a=>e({url:"/namespace/getUserNamespaces",method:"post",data:a}))(n).then((e=>{u.value=e.data.namespaces,i.value=e.data.count||50}))};p();const c=o(!1);let m=t({name:"",raft_id:""});const f=o([]);a(n).then((e=>{f.value=[];for(let a=0;a<e.data.cluster.length;a++)f.value.push(e.data.cluster[a])}));const h=o(!1);let _=t({username:"",auth:"",namespace_id:""});return{query:n,tableData:u,pageTotal:i,handleSearch:()=>{n.pageIndex=1,p()},handlePageChange:e=>{n.pageIndex=e,p()},goConfig:(e,a)=>{d.push({name:"config",query:{id:e,namespace:a}})},addVisible:c,form:m,options:f,handleAdd:()=>{c.value=!0},addNamespace:()=>{r.confirm("确定要添加吗？","提示",{type:"warning"}).then((()=>{(a=>e({url:"/namespace/newNamespace",method:"post",data:a}))({name:m.name,raft_id:m.raft_id}).then((e=>{0===e.status?(s.success("添加成功"),c.value=!1):s.error("添加失败")})),p()})).catch((()=>{}))},goCluster:e=>{d.push({name:"monitor",query:{raft_id:e}})},setAuthVisible:h,form2:_,setNamespaceAuth:()=>{r.confirm("确定要设置吗？","提示",{type:"warning"}).then((()=>{(a=>e({url:"/namespace/setAuth",method:"post",data:a}))({namespace_id:_.namespace_id,username:_.username,type:_.auth}).then((e=>{0===e.status?(s.success("设置成功"),c.value=!1):s.error("设置失败")})),p()})).catch((()=>{}))},setAuth:e=>{_.username="",_.auth="",_.namespace_id=e,h.value=!0}}}},V=y();d("data-v-01c13e58");const k={class:"crumbs"},C=c("i",{class:"el-icon-lx-cascades"},null,-1),w=m(" 命名空间 "),v={class:"container"},x={class:"handle-box"},A=m("搜索"),N=m("新增"),q={key:0},z={key:1},U={key:2},I=m("设置权限"),j=m("设置权限"),D={class:"pagination"},S={style:{float:"left"}},R={style:{float:"right",color:"#8492a6","font-size":"13px"}},T={class:"dialog-footer"},B=m("取 消"),P=m("确 定"),F=m("Normal"),L=m("Readonly"),O=m("Banned"),E={class:"dialog-footer"},G=m("取 消"),H=m("确 定");n();const J=V(((e,a,l,t,o,d)=>{const n=u("el-breadcrumb-item"),r=u("el-breadcrumb"),s=u("el-input"),y=u("el-button"),b=u("el-table-column"),J=u("el-tag"),K=u("el-table"),M=u("el-pagination"),Q=u("el-form-item"),W=u("el-option"),X=u("el-select"),Y=u("el-form"),Z=u("el-dialog");return i(),p("div",null,[c("div",k,[c(r,{separator:"/"},{default:V((()=>[c(n,null,{default:V((()=>[C,w])),_:1})])),_:1})]),c("div",v,[c("div",x,[c(s,{modelValue:t.query.name,"onUpdate:modelValue":a[1]||(a[1]=e=>t.query.name=e),placeholder:"命名空间",class:"handle-input mr10"},null,8,["modelValue"]),c(y,{type:"primary",icon:"el-icon-search",onClick:t.handleSearch},{default:V((()=>[A])),_:1},8,["onClick"]),c(y,{type:"primary",icon:"el-icon-plus",onClick:t.handleAdd},{default:V((()=>[N])),_:1},8,["onClick"])]),c(K,{data:t.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:V((()=>[c(b,{prop:"id",label:"ID",width:"55",align:"center"}),c(b,{prop:"name",label:"命名空间"},{default:V((e=>[c(y,{size:"medium",round:"",onClick:a=>t.goConfig(e.row.id,e.row.name)},{default:V((()=>[m(f(e.row.name),1)])),_:2},1032,["onClick"])])),_:1}),c(b,{prop:"username",label:"所有者"}),c(b,{prop:"raft_id",label:"集群ID"},{default:V((e=>[c(y,{size:"medium",round:"",onClick:a=>t.goCluster(e.row.raft_id)},{default:V((()=>[m(f(e.row.raft_id),1)])),_:2},1032,["onClick"])])),_:1}),c(b,{prop:"private_key",label:"密钥"}),c(b,{label:"状态",align:"center"},{default:V((e=>[c(J,{effect:"dark",type:0===e.row.type?"success":1===e.row.type?"warning":2===e.row.type?"danger":"info"},{default:V((()=>[0===e.row.type?(i(),p("p",q," Owner ")):1===e.row.type?(i(),p("p",z," Normal ")):2===e.row.type?(i(),p("p",U," Readonly ")):h("",!0)])),_:2},1032,["type"])])),_:1}),c(b,{label:"操作",align:"center"},{default:V((e=>[0===e.row.type?(i(),p(y,{key:0,type:"text",icon:"el-icon-s-grid",onClick:a=>t.setAuth(e.row.id)},{default:V((()=>[I])),_:2},1032,["onClick"])):(i(),p(y,{key:1,disabled:"",type:"text",icon:"el-icon-s-grid"},{default:V((()=>[j])),_:1}))])),_:1})])),_:1},8,["data"]),c("div",D,[c(M,{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal,onCurrentChange:t.handlePageChange},null,8,["current-page","page-size","total","onCurrentChange"])])]),c(Z,{title:"新建命名空间",modelValue:t.addVisible,"onUpdate:modelValue":a[5]||(a[5]=e=>t.addVisible=e),width:"30%"},{footer:V((()=>[c("span",T,[c(y,{onClick:a[4]||(a[4]=e=>t.addVisible=!1)},{default:V((()=>[B])),_:1}),c(y,{type:"primary",onClick:t.addNamespace},{default:V((()=>[P])),_:1},8,["onClick"])])])),default:V((()=>[c(Y,{"label-width":"70px"},{default:V((()=>[c(Q,{label:"命名"},{default:V((()=>[c(s,{modelValue:t.form.name,"onUpdate:modelValue":a[2]||(a[2]=e=>t.form.name=e)},null,8,["modelValue"])])),_:1}),c(Q,{label:"集群"},{default:V((()=>[c(X,{modelValue:t.form.raft_id,"onUpdate:modelValue":a[3]||(a[3]=e=>t.form.raft_id=e),placeholder:"请选择"},{default:V((()=>[(i(!0),p(_,null,g(t.options,(e=>(i(),p(W,{key:e.raft_id,label:e.raft_id,value:e.raft_id},{default:V((()=>[c("span",S,f(e.raft_id),1),c("span",R,f(e.address),1)])),_:2},1032,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"]),c(Z,{title:"设置权限",modelValue:t.setAuthVisible,"onUpdate:modelValue":a[9]||(a[9]=e=>t.setAuthVisible=e),width:"30%"},{footer:V((()=>[c("span",E,[c(y,{onClick:a[8]||(a[8]=e=>t.setAuthVisible=!1)},{default:V((()=>[G])),_:1}),c(y,{type:"primary",onClick:t.setNamespaceAuth},{default:V((()=>[H])),_:1},8,["onClick"])])])),default:V((()=>[c(Y,{"label-width":"70px"},{default:V((()=>[c(Q,{label:"用户名"},{default:V((()=>[c(s,{modelValue:t.form2.username,"onUpdate:modelValue":a[6]||(a[6]=e=>t.form2.username=e)},null,8,["modelValue"])])),_:1}),c(Q,{label:"权限"},{default:V((()=>[c(X,{modelValue:t.form2.auth,"onUpdate:modelValue":a[7]||(a[7]=e=>t.form2.auth=e),placeholder:"请选择"},{default:V((()=>[c(W,{label:"Normal",value:1,key:"1"},{default:V((()=>[c(J,{type:"warning"},{default:V((()=>[F])),_:1})])),_:1}),c(W,{label:"Readonly",value:2,key:"2"},{default:V((()=>[c(J,{type:"danger"},{default:V((()=>[L])),_:1})])),_:1}),c(W,{label:"Abandon",value:-1,key:"-1"},{default:V((()=>[c(J,{type:"info"},{default:V((()=>[O])),_:1})])),_:1})])),_:1},8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"])])}));b.render=J,b.__scopeId="data-v-01c13e58";export default b;
