"use strict";(globalThis["webpackChunkcontrol"]=globalThis["webpackChunkcontrol"]||[]).push([[97],{4097:(e,t,a)=>{a.r(t),a.d(t,{default:()=>oe});var l=a(9835),o=a(6970),r=a(499),s=a(6647),n=a(8339),i=a(2320),u=a(6397);const c={class:"row justify-start items-center"},d={class:"title-h4 top-bar-text"},m=(0,l.aZ)({__name:"TopBar",setup(e){const t=(0,s.QT)().t,a=(0,u.n)(),m=[5,20,40,60,120,180].map((e=>({label:e+"s",value:e}))),p=(0,l.Fl)((()=>{const e=(0,n.yj)();return"/overview"===e.path})),g=(0,l.Fl)((()=>{const e=(0,n.yj)();return e.meta.title?t(e.meta.title.toString()):""})),v=(0,l.Fl)((()=>{const e=(0,n.yj)();return e.meta.icon?e.meta.icon.toString():""})),f=(0,l.Fl)({get:()=>a.getRefreshInterval,set:e=>a.setRefreshInterval(e)});return(e,t)=>{const a=(0,l.up)("q-icon"),s=(0,l.up)("q-space"),n=(0,l.up)("q-toolbar"),u=(0,l.up)("q-header");return(0,l.wg)(),(0,l.j4)(u,{class:"bg-transparent q-pa-sm"},{default:(0,l.w5)((()=>[(0,l.Wm)(n,null,{default:(0,l.w5)((()=>[(0,l._)("div",c,[(0,l.Wm)(a,{color:"primary",name:v.value,size:"md",class:"q-mr-sm"},null,8,["name"]),(0,l._)("div",d,(0,o.zw)(g.value),1)]),(0,l.Wm)(s),p.value?((0,l.wg)(),(0,l.j4)(i.Z,{key:0,modelValue:f.value,"onUpdate:modelValue":t[0]||(t[0]=e=>f.value=e),options:(0,r.SU)(m),label:e.$t("topBar.selectRefreshRate"),class:"top-bar-refresh-rate-select"},null,8,["modelValue","options","label"])):(0,l.kq)("",!0)])),_:1})])),_:1})}}});var p=a(6602),g=a(1663),v=a(2857),f=a(136),b=a(9984),h=a.n(b);const w=m,y=w;h()(m,"components",{QHeader:p.Z,QToolbar:g.Z,QIcon:v.Z,QSpace:f.Z});var _=a(7747);const x={class:"flex justify-between"},q=(0,l.aZ)({__name:"ProfileCard",setup(e){function t(){const e=new _.Z;e.logout()}return(e,a)=>{const r=(0,l.up)("q-tooltip"),s=(0,l.up)("q-btn");return(0,l.wg)(),(0,l.iD)("div",x,[(0,l.Wm)(s,{dense:"",color:"primary",icon:"settings",size:"sm",to:"/settings"},{default:(0,l.w5)((()=>[(0,l.Wm)(r,{anchor:"bottom middle",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("profileCard.btnSettings")),1)])),_:1})])),_:1}),(0,l.Wm)(s,{dense:"",color:"grey-8",icon:"logout",size:"sm",onClick:a[0]||(a[0]=e=>t())},{default:(0,l.w5)((()=>[(0,l.Wm)(r,{anchor:"bottom middle",class:"bg-grey-8",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("profileCard.btnLogout")),1)])),_:1})])),_:1})])}}});var k=a(1639),I=a(8879),P=a(6858);const W=(0,k.Z)(q,[["__scopeId","data-v-4842b75a"]]),Z=W;h()(q,"components",{QBtn:I.Z,QTooltip:P.Z});var M=a(9302);const U={style:{"text-align":"center",padding:"0 8px"}},z=["src"],Q={class:"column items-center"},S={class:"col",style:{"font-size":"12px"}},$=(0,l.aZ)({__name:"SideBarMenu",setup(e){const t=(0,s.QT)().t,a={right:"4px",borderRadius:"5px",backgroundColor:"#c97350",width:"5px",opacity:.75},i={right:"2px",borderRadius:"9px",backgroundColor:"#c97350",width:"9px",opacity:.2},u=(0,r.iH)(),c=[{menuPosition:6,title:"Backups",icon:"backup",path:"/backups",disabled:!0,isMenuItem:!0},{menuPosition:7,title:"Images",icon:"image",path:"/images",disabled:!0,isMenuItem:!0},{menuPosition:8,title:"CI/CD",icon:"settings_suggest",path:"/ci-cd",disabled:!0,isMenuItem:!1},{menuPosition:9,title:"Security",icon:"security",path:"/security",disabled:!0,isMenuItem:!0},{menuPosition:10,title:"Metrics&Logs",icon:"receipt",path:"/metrics-logs",disabled:!0,isMenuItem:!0},{menuPosition:11,title:"Terminal",icon:"terminal",path:"/terminal",disabled:!0,isMenuItem:!0},{menuPosition:12,title:"IAM",icon:"group",path:"/iam",disabled:!0,isMenuItem:!0},{menuPosition:13,title:"Server Settings",icon:"settings",path:"/server-settings",disabled:!0,isMenuItem:!0}],d=(0,l.Fl)((()=>{const e=(0,M.Z)();return e.dark.isActive?"/_/assets/control-logo-dark.svg":"/_/assets/control-logo-light.svg"}));function m(){const e=(0,n.tv)().getRoutes();let a=[];e.forEach((e=>{const l=e.children||[];l.forEach((e=>{var l,o,r,s,n,i,u,c,d,m;!1!==(null===(l=e.meta)||void 0===l?void 0:l.isMenuItem)&&a.push({menuPosition:null!==(r=null===(o=e.meta)||void 0===o?void 0:o.menuPosition)&&void 0!==r?r:0,title:null!==(n=t(`${null===(s=e.meta)||void 0===s?void 0:s.title}`))&&void 0!==n?n:"",icon:null!==(c=null===(u=null===(i=e.meta)||void 0===i?void 0:i.icon)||void 0===u?void 0:u.toString())&&void 0!==c?c:"",path:e.path,disabled:null!==(m=null===(d=e.meta)||void 0===d?void 0:d.disabled)&&void 0!==m&&m})}))}));let l=a.concat(c);return l.sort(((e,t)=>e.menuPosition<t.menuPosition?-1:e.menuPosition>t.menuPosition?1:0))}return(0,l.wF)((()=>{u.value=m()})),(e,s)=>{const n=(0,l.up)("router-link"),c=(0,l.up)("q-avatar"),m=(0,l.up)("q-tooltip"),p=(0,l.up)("q-item"),g=(0,l.up)("q-list"),v=(0,l.up)("q-scroll-area"),f=(0,l.up)("q-drawer"),b=(0,l.Q2)("ripple");return(0,l.wg)(),(0,l.j4)(f,{mini:!0,side:"left",bordered:"","show-if-above":"","mini-width":100},{default:(0,l.w5)((()=>[(0,l._)("div",U,[(0,l.Wm)(n,{to:"/"},{default:(0,l.w5)((()=>[(0,l._)("img",{src:d.value,alt:"Speedia Control",style:{"margin-top":"20px",height:"1.82rem",width:"100%"}},null,8,z)])),_:1})]),(0,l.Wm)(Z,{class:"q-pt-md q-px-md"}),(0,l.Wm)(v,{"thumb-style":a,"bar-style":i,style:{height:"calc(100% - 100px)","margin-top":"10px"}},{default:(0,l.w5)((()=>[(0,l.Wm)(g,{style:{padding:"0 4px","text-align":"center","overflow-x":"hidden"}},{default:(0,l.w5)((()=>[((0,l.wg)(!0),(0,l.iD)(l.HY,null,(0,l.Ko)(u.value,((e,a)=>(0,l.wy)(((0,l.wg)(),(0,l.j4)(p,{key:a,to:e.path,disable:e.disabled,clickable:""},{default:(0,l.w5)((()=>[(0,l._)("div",Q,[(0,l.Wm)(c,{icon:e.icon,style:{"background-color":"rgba(0, 0, 0, 0.25)"}},null,8,["icon"]),(0,l._)("div",S,(0,o.zw)(e.title),1)]),e.disabled?((0,l.wg)(),(0,l.j4)(m,{key:0,anchor:"center end",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)((0,r.SU)(t)("sideBarMenu.disabled")),1)])),_:1})):(0,l.kq)("",!0)])),_:2},1032,["to","disable"])),[[b]]))),128))])),_:1})])),_:1})])),_:1})}}});var j=a(906),B=a(6663),C=a(3246),F=a(490),T=a(1357),R=a(1136);const L=$,A=L;h()($,"components",{QDrawer:j.Z,QScrollArea:B.Z,QList:C.Z,QItem:F.Z,QAvatar:T.Z,QTooltip:P.Z}),h()($,"directives",{Ripple:R.Z});var H=a(7178);const D={class:"flex justify-end items-center"},V={class:"absolute-full flex flex-center"},E={class:"absolute-full flex flex-center"},K={class:"absolute-full flex flex-center"},Y=(0,l.aZ)({__name:"FooterBar",setup(e){const t=new H.Z,a=(0,u.n)(),s=(0,r.iH)(),n=(0,r.iH)(),i=(0,l.Fl)((()=>a.getRefreshInterval));function c(e){return e<50?"green":e<80?"orange":"red"}function d(){s.value&&clearInterval(s.value),s.value=setInterval((()=>{t.getSystemInfo().then((e=>{n.value=e.data.body})).catch((e=>{console.error(e)}))}),1e3*i.value)}return(0,l.wF)((()=>{t.getSystemInfo().then((e=>{n.value=e.data.body})).catch((e=>{console.error(e)})).finally((()=>{d()}))})),(0,l.Ah)((()=>{s.value&&clearInterval(s.value)})),(e,t)=>{const a=(0,l.up)("q-tooltip"),r=(0,l.up)("q-icon"),s=(0,l.up)("q-badge"),i=(0,l.up)("q-linear-progress"),u=(0,l.up)("q-footer");return n.value?((0,l.wg)(),(0,l.j4)(u,{key:0,bordered:"",class:"bg-footer q-px-lg"},{default:(0,l.w5)((()=>[(0,l._)("div",D,[(0,l.Wm)(r,{name:"terminal",size:"1.618rem",class:"disabled q-mr-md"},{default:(0,l.w5)((()=>[(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("footerBar.disabled")),1)])),_:1})])),_:1}),(0,l.Wm)(r,{name:"speed",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(i,{stripe:"",rounded:"",size:"20px",class:"q-mr-md",value:Math.trunc(n.value.resourceUsage.cpuPercent)/100,color:c(Math.trunc(n.value.resourceUsage.cpuPercent)),label:`${Math.trunc(n.value.resourceUsage.cpuPercent)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",V,[(0,l.Wm)(s,{color:"white","text-color":"dark",label:`${Math.trunc(n.value.resourceUsage.cpuPercent)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("footerBar.cpuUsage",{cpuUsage:Math.trunc(n.value.resourceUsage.cpuPercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,l.Wm)(r,{name:"memory",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(i,{stripe:"",rounded:"",class:"q-mr-md",size:"20px",value:Math.trunc(n.value.resourceUsage.memoryPercent)/100,color:c(Math.trunc(n.value.resourceUsage.memoryPercent)),label:`${Math.trunc(n.value.resourceUsage.memoryPercent)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",E,[(0,l.Wm)(s,{color:"white","text-color":"dark",label:`${Math.trunc(n.value.resourceUsage.memoryPercent)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("footerBar.ramUsage",{ramUsage:Math.trunc(n.value.resourceUsage.memoryPercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,l.Wm)(r,{name:"sd_card",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(i,{stripe:"",rounded:"",size:"20px",value:Math.trunc(n.value.resourceUsage.storageInfo[0].usedPercent)/100,color:c(Math.trunc(n.value.resourceUsage.storageInfo[0].usedPercent)),label:`${Math.trunc(n.value.resourceUsage.storageInfo[0].usedPercent)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",K,[(0,l.Wm)(s,{color:"white","text-color":"dark",label:`${Math.trunc(n.value.resourceUsage.storageInfo[0].usedPercent)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,o.zw)(e.$t("footerBar.storageInfo",{storageInfo:Math.trunc(n.value.resourceUsage.storageInfo[0].usedPercent)})),1)])),_:1})])),_:1},8,["value","color","label"])])])),_:1})):(0,l.kq)("",!0)}}});var G=a(1378),J=a(8289),N=a(990);const O=(0,k.Z)(Y,[["__scopeId","data-v-202c37f2"]]),X=O;h()(Y,"components",{QFooter:G.Z,QIcon:v.Z,QTooltip:P.Z,QLinearProgress:J.Z,QBadge:N.Z});const ee=(0,l.aZ)({__name:"MainLayout",setup(e){return(e,t)=>{const a=(0,l.up)("router-view"),o=(0,l.up)("q-page-container"),r=(0,l.up)("q-layout");return(0,l.wg)(),(0,l.j4)(r,{view:"lhh Lpr lFf"},{default:(0,l.w5)((()=>[(0,l.Wm)(y),(0,l.Wm)(A),(0,l.Wm)(o,null,{default:(0,l.w5)((()=>[(0,l.Wm)(a)])),_:1}),(0,l.Wm)(X)])),_:1})}}});var te=a(249),ae=a(2133);const le=ee,oe=le;h()(ee,"components",{QLayout:te.Z,QPageContainer:ae.Z})}}]);