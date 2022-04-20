import{s as e}from"./request.a57216d8.js";import{a}from"./cluster.e0a8b9ec.js";import{d as l,C as t,B as d,r as o,o as n,e as r,f as s,g as u,w as i,i as p,t as c,h as m,c as f,F as h,l as _,L as g,D as y,p as b,j as V}from"./vendor.da3fc754.js";import{_ as k}from"./index.0cb5857a.js";const C={name:"namespace",setup(){const o=l(),n=t({name:"",pageIndex:1,pageSize:10}),r=d([]),s=d(0),u=()=>{(a=>e({url:"/namespace/getUserNamespaces",method:"post",data:a}))(n).then((e=>{r.value=e.data.namespaces,s.value=e.data.count||50}))};u();const i=d(!1);let p=t({name:"",raft_id:""});const c=d([]);a(n).then((e=>{c.value=[];for(let a=0;a<e.data.cluster.length;a++)c.value.push(e.data.cluster[a])}));const m=d(!1);let f=t({username:"",auth:"",namespace_id:""});return{query:n,tableData:r,pageTotal:s,handleSearch:()=>{n.pageIndex=1,u()},handlePageChange:e=>{n.pageIndex=e,u()},goConfig:(e,a)=>{o.push({name:"config",query:{id:e,namespace:a}})},addVisible:i,form:p,options:c,handleAdd:()=>{i.value=!0},addNamespace:()=>{g.confirm("确定要添加吗？","提示",{type:"warning"}).then((()=>{(a=>e({url:"/namespace/newNamespace",method:"post",data:a}))({name:p.name,raft_id:p.raft_id}).then((e=>{0===e.status?(y.success("添加成功"),i.value=!1):y.error("添加失败")})),u()})).catch((()=>{}))},goCluster:()=>{},setAuthVisible:m,form2:f,setNamespaceAuth:()=>{g.confirm("确定要设置吗？","提示",{type:"warning"}).then((()=>{(a=>e({url:"/namespace/setAuth",method:"post",data:a}))({namespace_id:f.namespace_id,username:f.username,type:f.auth}).then((e=>{0===e.status?(y.success("设置成功"),i.value=!1):y.error("设置失败")})),u()})).catch((()=>{}))},setAuth:e=>{f.username="",f.auth="",f.namespace_id=e,m.value=!0}}}},w={class:"crumbs"},v=(e=>(b("data-v-39893386"),e=e(),V(),e))((()=>s("i",{class:"el-icon-lx-cascades"},null,-1))),x=p(" 命名空间 "),A={class:"container"},N={class:"handle-box"},U=p("搜索"),q=p("新增"),z={key:0},I={key:1},j={key:2},D=p("设置权限"),S=p("设置权限"),R={class:"pagination"},T={style:{float:"left"}},B={style:{float:"right",color:"#8492a6","font-size":"13px"}},P={class:"dialog-footer"},F=p("取 消"),L=p("确 定"),O=p("Normal"),E=p("Readonly"),G=p("Banned"),H={class:"dialog-footer"},J=p("取 消"),K=p("确 定");var M=k(C,[["render",function(e,a,l,t,d,g){const y=o("el-breadcrumb-item"),b=o("el-breadcrumb"),V=o("el-input"),k=o("el-button"),C=o("el-table-column"),M=o("el-tag"),Q=o("el-table"),W=o("el-pagination"),X=o("el-form-item"),Y=o("el-option"),Z=o("el-select"),$=o("el-form"),ee=o("el-dialog");return n(),r("div",null,[s("div",w,[u(b,{separator:"/"},{default:i((()=>[u(y,null,{default:i((()=>[v,x])),_:1})])),_:1})]),s("div",A,[s("div",N,[u(V,{modelValue:t.query.name,"onUpdate:modelValue":a[0]||(a[0]=e=>t.query.name=e),placeholder:"命名空间",class:"handle-input mr10"},null,8,["modelValue"]),u(k,{type:"primary",icon:"el-icon-search",onClick:t.handleSearch},{default:i((()=>[U])),_:1},8,["onClick"]),u(k,{type:"primary",icon:"el-icon-plus",onClick:t.handleAdd},{default:i((()=>[q])),_:1},8,["onClick"])]),u(Q,{data:t.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:i((()=>[u(C,{prop:"id",label:"ID",width:"55",align:"center"}),u(C,{prop:"name",label:"命名空间"},{default:i((e=>[u(k,{size:"medium",round:"",onClick:a=>t.goConfig(e.row.id,e.row.name)},{default:i((()=>[p(c(e.row.name),1)])),_:2},1032,["onClick"])])),_:1}),u(C,{prop:"username",label:"所有者"}),u(C,{prop:"raft_id",label:"集群ID"},{default:i((e=>[u(k,{size:"medium",round:"",onClick:a=>t.goCluster(e.row.raft_id)},{default:i((()=>[p(c(e.row.raft_id),1)])),_:2},1032,["onClick"])])),_:1}),u(C,{prop:"private_key",label:"密钥"}),u(C,{label:"状态",align:"center"},{default:i((e=>[u(M,{effect:"dark",type:0===e.row.type?"success":1===e.row.type?"warning":2===e.row.type?"danger":"info"},{default:i((()=>[0===e.row.type?(n(),r("p",z," Owner ")):1===e.row.type?(n(),r("p",I," Normal ")):2===e.row.type?(n(),r("p",j," Readonly ")):m("",!0)])),_:2},1032,["type"])])),_:1}),u(C,{label:"操作",align:"center"},{default:i((e=>[0===e.row.type?(n(),f(k,{key:0,type:"text",icon:"el-icon-s-grid",onClick:a=>t.setAuth(e.row.id)},{default:i((()=>[D])),_:2},1032,["onClick"])):(n(),f(k,{key:1,disabled:"",type:"text",icon:"el-icon-s-grid"},{default:i((()=>[S])),_:1}))])),_:1})])),_:1},8,["data"]),s("div",R,[u(W,{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal,onCurrentChange:t.handlePageChange},null,8,["current-page","page-size","total","onCurrentChange"])])]),u(ee,{title:"新建命名空间",modelValue:t.addVisible,"onUpdate:modelValue":a[4]||(a[4]=e=>t.addVisible=e),width:"30%"},{footer:i((()=>[s("span",P,[u(k,{onClick:a[3]||(a[3]=e=>t.addVisible=!1)},{default:i((()=>[F])),_:1}),u(k,{type:"primary",onClick:t.addNamespace},{default:i((()=>[L])),_:1},8,["onClick"])])])),default:i((()=>[u($,{"label-width":"70px"},{default:i((()=>[u(X,{label:"命名"},{default:i((()=>[u(V,{modelValue:t.form.name,"onUpdate:modelValue":a[1]||(a[1]=e=>t.form.name=e)},null,8,["modelValue"])])),_:1}),u(X,{label:"集群"},{default:i((()=>[u(Z,{modelValue:t.form.raft_id,"onUpdate:modelValue":a[2]||(a[2]=e=>t.form.raft_id=e),placeholder:"请选择"},{default:i((()=>[(n(!0),r(h,null,_(t.options,(e=>(n(),f(Y,{key:e.raft_id,label:e.raft_id,value:e.raft_id},{default:i((()=>[s("span",T,c(e.raft_id),1),s("span",B,c(e.address),1)])),_:2},1032,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"]),u(ee,{title:"设置权限",modelValue:t.setAuthVisible,"onUpdate:modelValue":a[8]||(a[8]=e=>t.setAuthVisible=e),width:"30%"},{footer:i((()=>[s("span",H,[u(k,{onClick:a[7]||(a[7]=e=>t.setAuthVisible=!1)},{default:i((()=>[J])),_:1}),u(k,{type:"primary",onClick:t.setNamespaceAuth},{default:i((()=>[K])),_:1},8,["onClick"])])])),default:i((()=>[u($,{"label-width":"70px"},{default:i((()=>[u(X,{label:"用户名"},{default:i((()=>[u(V,{modelValue:t.form2.username,"onUpdate:modelValue":a[5]||(a[5]=e=>t.form2.username=e)},null,8,["modelValue"])])),_:1}),u(X,{label:"权限"},{default:i((()=>[u(Z,{modelValue:t.form2.auth,"onUpdate:modelValue":a[6]||(a[6]=e=>t.form2.auth=e),placeholder:"请选择"},{default:i((()=>[u(Y,{label:"Normal",value:1,key:"1"},{default:i((()=>[u(M,{type:"warning"},{default:i((()=>[O])),_:1})])),_:1}),u(Y,{label:"Readonly",value:2,key:"2"},{default:i((()=>[u(M,{type:"danger"},{default:i((()=>[E])),_:1})])),_:1}),u(Y,{label:"Abandon",value:-1,key:"-1"},{default:i((()=>[u(M,{type:"info"},{default:i((()=>[G])),_:1})])),_:1})])),_:1},8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"])])}],["__scopeId","data-v-39893386"]]);export{M as default};
