import{f as e,B as a,A as s,u as r,p as o,b as l,C as t,r as m,o as u,c as d,g as n,E as p,w as i,h as c}from"./vendor.971a53c8.js";import{u as g}from"./index.d0a782d5.js";const f={setup(){const o=e(),l=a({username:"lxh001",password:"123456"}),m=s(null);return r().commit("clearTags"),{param:l,rules:{username:[{required:!0,message:"请输入用户名",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}]},login:m,submitForm:()=>{g(l).then((e=>{if(0!==e.status)return t.error("登录失败"),!1;t.success("登录成功"),localStorage.setItem("ms_username",l.username),localStorage.setItem("is_admin",e.data),o.push("/")}))}}}},w=i();o("data-v-67956cc0");const _={class:"login-wrap"},b={class:"ms-login"},h=n("div",{class:"ms-title"},"配置管理系统",-1),v={class:"login-btn"},V=c("登录");l();const x=w(((e,a,s,r,o,l)=>{const t=m("el-button"),i=m("el-input"),c=m("el-form-item"),g=m("el-form");return u(),d("div",_,[n("div",b,[h,n(g,{model:r.param,rules:r.rules,ref:"login","label-width":"0px",class:"ms-content"},{default:w((()=>[n(c,{prop:"username"},{default:w((()=>[n(i,{modelValue:r.param.username,"onUpdate:modelValue":a[1]||(a[1]=e=>r.param.username=e),placeholder:"username"},{prepend:w((()=>[n(t,{icon:"el-icon-user"})])),_:1},8,["modelValue"])])),_:1}),n(c,{prop:"password"},{default:w((()=>[n(i,{type:"password",placeholder:"password",modelValue:r.param.password,"onUpdate:modelValue":a[2]||(a[2]=e=>r.param.password=e),onKeyup:a[3]||(a[3]=p((e=>r.submitForm()),["enter"]))},{prepend:w((()=>[n(t,{icon:"el-icon-lock"})])),_:1},8,["modelValue"])])),_:1}),n("div",v,[n(t,{type:"primary",onClick:a[4]||(a[4]=e=>r.submitForm())},{default:w((()=>[V])),_:1})])])),_:1},8,["model","rules"])])])}));f.render=x,f.__scopeId="data-v-67956cc0";export default f;