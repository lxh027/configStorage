import"./index.cb2b8693.js";import{g as a,a as e}from"./cluster.192a7c6a.js";import{f as t,A as l,p as s,b as r,r as c,o as d,c as o,g as n,w as u,h as i}from"./vendor.971a53c8.js";const b={name:"dashboard",setup(){const s=localStorage.getItem("is_admin"),r=t(),c=l([]);"1"===s?a().then((a=>{c.value=a.data.cluster})):e({}).then((a=>{c.value=a.data.cluster}));return{tableData:c,goCluster:a=>{r.push({name:"monitor",query:{raft_id:a}})}}}},m=u();s("data-v-6c0ea77e");const p={class:"crumbs"},f=n("i",{class:"el-icon-lx-cascades"},null,-1),_=i(" 集群列表 "),g={class:"container"},v=i("查看监控");r();const h=m(((a,e,t,l,s,r)=>{const u=c("el-breadcrumb-item"),i=c("el-breadcrumb"),b=c("el-table-column"),h=c("el-button"),x=c("el-table");return d(),o("div",null,[n("div",p,[n(i,{separator:"/"},{default:m((()=>[n(u,null,{default:m((()=>[f,_])),_:1})])),_:1})]),n("div",g,[n(x,{data:l.tableData,border:"",class:"table",ref:"multipleTable","header-cell-class-name":"table-header"},{default:m((()=>[n(b,{prop:"raft_id",label:"集群ID"}),n(b,{prop:"address",label:"地址"}),n(b,{label:"操作",align:"center"},{default:m((a=>[n(h,{type:"text",icon:"el-icon-s-grid",onClick:e=>l.goCluster(a.row.raft_id)},{default:m((()=>[v])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])])])}));b.render=h,b.__scopeId="data-v-6c0ea77e";export default b;
