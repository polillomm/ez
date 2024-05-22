"use strict";(globalThis["webpackChunkcontrol"]=globalThis["webpackChunkcontrol"]||[]).push([[762],{4806:(e,t,a)=>{a.r(t),a.d(t,{default:()=>yt});var l=a(9835),n=a(499),i=a(6970),o=a(1809);const s=(0,o.Q_)("mappings",{state:()=>({keyMappingsTable:0,mappingsList:[],selectedMapping:{},selectedTarget:{},showCreateMappingDialog:!1,showDeleteMappingDialog:!1,showCreateTargetDialog:!1,showDeleteTargetDialog:!1}),getters:{getKeyMappingsTable(e){return e.keyMappingsTable},getMappingsList(e){return e.mappingsList},getSelectedMapping(e){return e.selectedMapping},getSelectedTarget(e){return e.selectedTarget},getShowCreateMappingDialog(e){return e.showCreateMappingDialog},getShowDeleteMappingDialog(e){return e.showDeleteMappingDialog},getShowCreateTargetDialog(e){return e.showCreateTargetDialog},getShowDeleteTargetDialog(e){return e.showDeleteTargetDialog}},actions:{setKeyMappingsTable(e){this.keyMappingsTable=e},setMappingsList(e){this.mappingsList=e},setSelectedMapping(e){this.selectedMapping=e},setSelectedTarget(e){this.selectedTarget=e},setShowCreateMappingDialog(e){this.showCreateMappingDialog=e},setShowDeleteMappingDialog(e){this.showDeleteMappingDialog=e},setShowCreateTargetDialog(e){this.showCreateTargetDialog=e},setShowDeleteTargetDialog(e){this.showDeleteTargetDialog=e}}}),p={class:"flex justify-center items-center"},r=(0,l.aZ)({__name:"MappingsTableActionsTd",props:{mapping:{}},setup(e){const t=e;function a(){const e=s();e.setSelectedMapping(t.mapping),e.setShowDeleteMappingDialog(!0)}return(e,t)=>{const n=(0,l.up)("q-tooltip"),o=(0,l.up)("q-btn");return(0,l.wg)(),(0,l.iD)("div",p,[(0,l.Wm)(o,{color:"negative",icon:"delete",size:"md",dense:"",outline:"",class:"q-mx-xs",onClick:t[0]||(t[0]=e=>a())},{default:(0,l.w5)((()=>[(0,l.Wm)(n,{class:"bg-negative text-white",style:{"font-size":"14px"},offset:[10,10]},{default:(0,l.w5)((()=>[(0,l.Uk)((0,i.zw)(e.$t("mappingsTableActionsTd.deleteBtn")),1)])),_:1})])),_:1})])}}});var c=a(8879),u=a(6858),g=a(9984),d=a.n(g);const m=r,v=m;d()(r,"components",{QBtn:c.Z,QTooltip:u.Z});var b=a(9573);const w={class:"flex justify-between q-mx-xs"},h={class:"text-h6 mappings-card-title"},f={class:"column",style:{width:"40%"}},D=(0,l.aZ)({__name:"MappingsTableMainTd",props:{mapping:{}},setup(e){return(e,t)=>{const a=(0,l.up)("q-card-section"),n=(0,l.up)("q-card");return(0,l.wg)(),(0,l.j4)(n,{flat:"",class:"q-pa-sm mappings-card"},{default:(0,l.w5)((()=>[(0,l._)("div",w,[(0,l._)("div",h," #"+(0,i.zw)(e.mapping.id)+" - "+(0,i.zw)(e.mapping.hostname?e.mapping.hostname:"*"),1),(0,l.Wm)(v,{mapping:e.mapping,class:"absolute-top-right q-ma-sm"},null,8,["mapping"])]),(0,l.Wm)(a,{class:"card-section-mappings"},{default:(0,l.w5)((()=>[(0,l._)("div",f,[(0,l.Wm)(b.Z,{icon:"person",label:e.mapping.accountName},null,8,["label"]),(0,l.Wm)(b.Z,{icon:"cable",label:e.mapping.protocol},null,8,["label"]),(0,l.Wm)(b.Z,{icon:"lan",label:e.mapping.publicPort.toString()},null,8,["label"])])])),_:1})])),_:1})}}});var y=a(4458),_=a(3190);const T=D,x=T;d()(D,"components",{QCard:y.Z,QCardSection:_.Z});const M={class:"default-chip-text q-px-xs"},Z=(0,l.aZ)({__name:"CustomChip",props:{labelLeft:{},labelRight:{},color:{default:"#c97350"},textColor:{default:"#fff"},fontSize:{default:"1.2499rem"}},setup(e){const t=e,a=(0,n.iH)({}),o=(0,n.iH)({});return(0,l.bv)((()=>{a.value={backgroundColor:t.color,color:t.textColor},o.value={border:"1px solid "+t.color,borderRadius:"5px",fontSize:t.fontSize}})),(e,t)=>((0,l.wg)(),(0,l.iD)("div",{style:(0,i.j5)(o.value)},[(0,l._)("div",{class:"q-px-xs float-left",style:(0,i.j5)(a.value)},(0,i.zw)(e.labelLeft),5),(0,l._)("span",M,(0,i.zw)(e.labelRight),1)],4))}}),C=Z,q=C,k={class:"flex justify-between items-center"},W={class:"row justify-start items-center"},S={class:"flex justify-end q-px-sm q-py-sm"},P=(0,l.aZ)({__name:"MappingsTableTargetTd",props:{mapping:{}},setup(e){const t=e,a=s(),n=(0,l.Fl)({get:()=>a.getShowDeleteTargetDialog,set:e=>{a.setShowDeleteMappingDialog(e)}}),o=(0,l.Fl)({get:()=>a.getSelectedTarget,set:e=>{a.setSelectedTarget(e)}});function p(e,a){o.value={mappingId:t.mapping.id,id:e,containerId:a},n.value=!0}function r(){a.setSelectedMapping(t.mapping),a.setShowCreateTargetDialog(!0)}return(e,t)=>{const a=(0,l.up)("q-tooltip"),n=(0,l.up)("q-btn"),o=(0,l.up)("q-card");return(0,l.wg)(),(0,l.iD)("div",null,[((0,l.wg)(!0),(0,l.iD)(l.HY,null,(0,l.Ko)(e.mapping.targets,((t,s)=>((0,l.wg)(),(0,l.iD)("div",{key:s},[(0,l.Wm)(o,{flat:"",class:"q-ma-sm q-py-sm q-px-sm target-td-card-bg"},{default:(0,l.w5)((()=>[(0,l._)("div",k,[(0,l._)("div",W,[(0,l.Wm)(q,{labelLeft:e.$t("mappingsTableTargetTd.containerHostname"),labelRight:t.containerHostname,color:"#39594E",fontSize:"0.7281rem"},null,8,["labelLeft","labelRight"])]),(0,l.Wm)(n,{color:"negative",icon:"delete",size:"md",dense:"",onClick:e=>p(t.id,t.containerId)},{default:(0,l.w5)((()=>[(0,l.Wm)(a,{class:"bg-negative text-white",style:{"font-size":"14px"},offset:[10,10]},{default:(0,l.w5)((()=>[(0,l.Uk)((0,i.zw)(e.$t("mappingsTableTargetTd.deleteBtn")),1)])),_:1})])),_:2},1032,["onClick"])])])),_:2},1024)])))),128)),(0,l._)("div",S,[(0,l.Wm)(n,{color:"primary",icon:"add",class:"full-width",outline:"",dense:"",onClick:t[0]||(t[0]=e=>r())},{default:(0,l.w5)((()=>[(0,l.Wm)(a,{class:"bg-primary text-white",style:{"font-size":"14px"},offset:[10,10]},{default:(0,l.w5)((()=>[(0,l.Uk)((0,i.zw)(e.$t("mappingsTableTargetTd.createTargetBtn")),1)])),_:1})])),_:1})])])}}}),I=P,$=I;d()(P,"components",{QCard:y.Z,QBtn:c.Z,QTooltip:u.Z});var H=a(5786);const Q={class:"flex justify-between itens-center full-width q-mb-sm"},L={class:"row justify-center q-mt-md"},j=(0,l.aZ)({__name:"MappingsTable",props:{mappingsList:{}},setup(e){const t=e,a=s(),i=(0,n.iH)(""),o=(0,n.iH)([{name:"hostname",align:"left",field:"hostname"},{name:"accountName",align:"left",field:"accountName"},{name:"protocol",align:"left",field:"protocol"},{name:"publicPort",align:"left",field:"publicPort"}]),p=(0,n.iH)({sortBy:"desc",descending:!1,page:1,rowsPerPage:10}),r=(0,l.Fl)({get:()=>a.getShowCreateMappingDialog,set:e=>{a.showCreateMappingDialog=e}}),c=(0,l.Fl)((()=>Math.ceil(t.mappingsList.length/p.value.rowsPerPage)));function u(){r.value=!0}function g(){a.showCreateTargetDialog=!0}return(e,t)=>{const a=(0,l.up)("q-icon"),n=(0,l.up)("q-input"),s=(0,l.up)("q-td"),r=(0,l.up)("q-tr"),d=(0,l.up)("q-table"),m=(0,l.up)("q-pagination");return(0,l.wg)(),(0,l.iD)(l.HY,null,[(0,l._)("div",Q,[(0,l.Wm)(n,{filled:"",debounce:"300",color:"primary",modelValue:i.value,"onUpdate:modelValue":t[0]||(t[0]=e=>i.value=e),label:e.$t("mappingsTable.searchInput"),style:{width:"300px"}},{prepend:(0,l.w5)((()=>[(0,l.Wm)(a,{name:"search"})])),_:1},8,["modelValue","label"]),(0,l._)("div",null,[(0,l.Wm)(H.Z,{label:e.$t("mappingsTable.createTargetBtn"),color:"primary",icon:"add",class:"q-mr-md",onClick:t[1]||(t[1]=e=>g())},null,8,["label"]),(0,l.Wm)(H.Z,{label:e.$t("mappingsTable.createMappingBtn"),color:"primary",icon:"add",onClick:t[2]||(t[2]=e=>u())},null,8,["label"])])]),(0,l.Wm)(d,{rows:e.mappingsList,columns:o.value,filter:i.value,pagination:p.value,"onUpdate:pagination":t[3]||(t[3]=e=>p.value=e),"no-data-label":e.$t("mappingsTable.noDataLabel"),"row-key":"hostname",color:"primary",flat:"",bordered:"","hide-header":"","hide-pagination":""},{body:(0,l.w5)((e=>[(0,l.Wm)(r,{props:e,"no-hover":""},{default:(0,l.w5)((()=>[(0,l.Wm)(s,{class:"mappings-td-main"},{default:(0,l.w5)((()=>[(0,l.Wm)(x,{mapping:e.row},null,8,["mapping"])])),_:2},1024),(0,l.Wm)(s,null,{default:(0,l.w5)((()=>[(0,l.Wm)($,{mapping:e.row},null,8,["mapping"])])),_:2},1024)])),_:2},1032,["props"])])),_:1},8,["rows","columns","filter","pagination","no-data-label"]),(0,l._)("div",L,[(0,l.Wm)(m,{modelValue:p.value.page,"onUpdate:modelValue":t[4]||(t[4]=e=>p.value.page=e),color:"primary",max:c.value,size:"md"},null,8,["modelValue","max"])])],64)}}});var z=a(6611),B=a(2857),F=a(422),U=a(1233),A=a(7220),V=a(996);const Y=j,K=Y;d()(j,"components",{QInput:z.Z,QIcon:B.Z,QTable:F.Z,QTr:U.Z,QTd:A.Z,QPagination:V.Z});const R={class:"q-pa-md"},N={class:"text-left",style:{width:"150px"}},E=(0,l._)("th",null,null,-1),X={class:"text-left"},G={class:"text-right"},J={class:"text-right"},O={class:"text-right"},ee={class:"text-right"},te={class:"text-right"};function ae(e,t){const a=(0,l.up)("q-skeleton"),n=(0,l.up)("q-markup-table");return(0,l.wg)(),(0,l.iD)("div",R,[(0,l.Wm)(n,null,{default:(0,l.w5)((()=>[(0,l._)("thead",null,[(0,l._)("tr",null,[(0,l._)("th",N,[(0,l.Wm)(a,{animation:"blink",type:"text"})]),((0,l.wg)(),(0,l.iD)(l.HY,null,(0,l.Ko)(4,(e=>(0,l._)("th",{key:e,class:"text-right"},[(0,l.Wm)(a,{animation:"blink",type:"text"})]))),64)),E])]),(0,l._)("tbody",null,[((0,l.wg)(),(0,l.iD)(l.HY,null,(0,l.Ko)(10,(e=>(0,l._)("tr",{key:e},[(0,l._)("td",X,[(0,l.Wm)(a,{animation:"blink",type:"text",width:"85px"})]),(0,l._)("td",G,[(0,l.Wm)(a,{animation:"blink",type:"text",width:"50px"})]),(0,l._)("td",J,[(0,l.Wm)(a,{animation:"blink",type:"text",width:"35px"})]),(0,l._)("td",O,[(0,l.Wm)(a,{animation:"blink",type:"text",width:"65px"})]),(0,l._)("td",ee,[(0,l.Wm)(a,{animation:"blink",type:"text",width:"25px"})]),(0,l._)("td",te,[(0,l.Wm)(a,{size:"25px"})])]))),64))])])),_:1})])}var le=a(1639),ne=a(6933),ie=a(7133);const oe={},se=(0,le.Z)(oe,[["render",ae]]),pe=se;d()(oe,"components",{QMarkupTable:ne.Z,QSkeleton:ie.ZP});var re=a(2320),ce=a(5971),ue=a(503);const ge=(0,l.aZ)({__name:"MappingsAccountSelect",emits:["update:selectedAccount"],setup(e,{emit:t}){const a=t,i=(0,n.iH)(!1),o=(0,n.iH)([]),s=(0,n.iH)(0);function p(){i.value=!0;const e=new ce.Z;e.getAccounts().then((e=>{if(o.value=[],0===e.data.body.length)return;o.value=r(e.data.body);const t=new ue.Z;s.value=t.getAccountId()?t.getAccountId():o.value[0].value})).catch((e=>{console.error(e)})).finally((()=>{i.value=!1}))}function r(e){return e.map((e=>({label:e.username,value:e.id})))}return(0,l.bv)((()=>{p()})),(0,l.YP)(s,(()=>{a("update:selectedAccount",s.value)})),(e,t)=>{const a=(0,l.up)("q-skeleton");return!0===i.value?((0,l.wg)(),(0,l.j4)(a,{key:0,animation:"wave",height:"75px"})):((0,l.wg)(),(0,l.j4)(re.Z,{key:1,selected:s.value,"onUpdate:selected":t[0]||(t[0]=e=>s.value=e),label:e.$t("mappingsAccountSelect.selectAccount"),options:o.value,dense:!1},null,8,["selected","label","options"]))}}}),de=ge,me=de;d()(ge,"components",{QSkeleton:ie.ZP});var ve=a(4304);const be=(0,l.aZ)({__name:"MappingsHostnameInput",props:{label:{default:""},prefix:{default:""},disable:{type:Boolean}},emits:["update:hostname","update:isValidHostname"],setup(e,{emit:t}){const a=e,i=t,o=(0,n.iH)(""),s=(0,l.Fl)((()=>{const e=/^[\p{L}\d][\p{L}\d\.]{0,1000}[\p{L}\d]$/u;return e.test(a.prefix+o.value.toString())}));return(0,l.YP)(o,(e=>{i("update:hostname",e)}),{immediate:!0}),(0,l.YP)(s,(e=>{i("update:isValidHostname",e)}),{immediate:!0}),(e,t)=>((0,l.wg)(),(0,l.j4)(ve.Z,{value:o.value,"onUpdate:value":t[0]||(t[0]=e=>o.value=e),label:e.label,prefix:e.prefix,disable:e.disable,rules:[()=>""!==o.value||e.$t("mappingsHostnameInput.hostnameIsRequired"),()=>s.value||e.$t("mappingsHostnameInput.hostnameIsInvalid")]},null,8,["value","label","prefix","disable","rules"]))}}),we=be,he=we,fe=(0,l.aZ)({__name:"MappingsProtocolSelect",emits:["update:selectedProtocol"],setup(e,{emit:t}){const a=t,i=["http","https","ws","wss","grpc","grpcs","tcp","udp"],o=(0,n.iH)("http");return(0,l.YP)(o,(e=>{a("update:selectedProtocol",e)}),{immediate:!0}),(e,t)=>((0,l.wg)(),(0,l.j4)(re.Z,{selected:o.value,"onUpdate:selected":t[0]||(t[0]=e=>o.value=e),options:i,label:e.$t("mappingsProtocolSelect.protocolSelect"),dense:!1},null,8,["selected","label"]))}}),De=fe,ye=De,_e=(0,l.aZ)({__name:"MappingsPublicPortInput",emits:["update:publicPort","update:isValidPublicPort"],setup(e,{emit:t}){const a=t,i=/^(\d{1,5})$/,o=(0,n.iH)(""),s=(0,l.Fl)((()=>!(""===o.value||parseInt(o.value)>65535)&&i.test(o.value)));return(0,l.YP)(o,(()=>{a("update:publicPort",o.value)})),(0,l.YP)(s,(()=>{a("update:isValidPublicPort",s.value)}),{immediate:!0}),(e,t)=>((0,l.wg)(),(0,l.j4)(ve.Z,{value:o.value,"onUpdate:value":t[0]||(t[0]=e=>o.value=e),label:e.$t("mappingsPublicPortInput.publicPortLabel"),maxLength:"5",rules:[()=>""!==o.value||e.$t("mappingsPublicPortInput.publicPortIsRequired"),()=>s.value||e.$t("mappingsPublicPortInput.publicPortIsInvalid")]},null,8,["value","label","rules"]))}}),Te=_e,xe=Te;var Me=a(9036),Ze=function(e,t,a,l){function n(e){return e instanceof a?e:new a((function(t){t(e)}))}return new(a||(a=Promise))((function(a,i){function o(e){try{p(l.next(e))}catch(t){i(t)}}function s(e){try{p(l["throw"](e))}catch(t){i(t)}}function p(e){e.done?a(e.value):n(e.value).then(o,s)}p((l=l.apply(e,t||[])).next())}))};class Ce extends Me.Z{getMappings(){return Ze(this,void 0,void 0,(function*(){return this.request.get("v1/mapping/")}))}createMapping(e){return Ze(this,void 0,void 0,(function*(){return this.request.post("v1/mapping/",e)}))}deleteMapping(e){return Ze(this,void 0,void 0,(function*(){return this.request.delete(`v1/mapping/${e.mappingId}/`)}))}createTarget(e){return Ze(this,void 0,void 0,(function*(){return this.request.post("v1/mapping/target/",e)}))}deleteTarget(e){return Ze(this,void 0,void 0,(function*(){return this.request.delete(`v1/mapping/${e.mappingId}/target/${e.targetId}/`)}))}}var qe=a(8900),ke=a(6647),We=a(5273);const Se={class:"flex justify-between items-center"},Pe={class:"title-dialog"},Ie={class:"flex justify-between"},$e=(0,l.aZ)({__name:"MappingsCreateDialog",setup(e){const t=(0,ke.QT)().t,a=s(),o=(0,n.iH)(""),p=(0,n.iH)(""),r=(0,n.iH)(""),c=(0,n.iH)(0),u=(0,n.iH)(!1),g=(0,n.iH)(!1),d=(0,l.Fl)({get:()=>a.getShowCreateMappingDialog,set:e=>a.setShowCreateMappingDialog(e)}),m=(0,l.Fl)({get:()=>a.getKeyMappingsTable,set:e=>a.setKeyMappingsTable(e)}),v=(0,l.Fl)((()=>u.value&&g.value&&0!==c.value));function b(){d.value=!1}function w(){(0,We.Q)();const e=new Ce;let a={hostname:o.value,accountId:c.value,publicPort:parseInt(p.value),protocol:r.value,containerIds:[]};e.createMapping(a).then((()=>{(0,qe.LX)(t("mappingsCreateDialog.createMappingSuccessfully")),m.value++,b()})).catch((e=>{console.error(e),(0,qe.s9)(e.response.data,t("mappingsCreateDialog.createMappingError"))})).finally((()=>{(0,We.Z)()}))}return(e,t)=>{const a=(0,l.up)("q-btn"),n=(0,l.up)("q-card-section"),s=(0,l.up)("q-card"),m=(0,l.up)("q-expansion-item"),h=(0,l.up)("q-list"),f=(0,l.up)("q-card-actions"),D=(0,l.up)("q-dialog");return(0,l.wg)(),(0,l.j4)(D,{modelValue:d.value,"onUpdate:modelValue":t[9]||(t[9]=e=>d.value=e),persistent:""},{default:(0,l.w5)((()=>[(0,l.Wm)(s,{style:{width:"700px","max-width":"80vw"},class:"dialog-card-bg"},{default:(0,l.w5)((()=>[(0,l._)("div",Se,[(0,l._)("div",Pe,(0,i.zw)(e.$t("mappingsCreateDialog.title")),1),(0,l.Wm)(a,{flat:"",round:"",dense:"",icon:"close",onClick:t[0]||(t[0]=e=>b())})]),(0,l.Wm)(n,{class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(he,{label:e.$t("mappingsCreateDialog.hostnameLabel"),"onUpdate:hostname":t[1]||(t[1]=e=>o.value=e),"onUpdate:isValidHostname":t[2]||(t[2]=e=>g.value=e)},null,8,["label"])])),_:1}),(0,l.Wm)(n,{class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l._)("div",Ie,[(0,l.Wm)(xe,{"onUpdate:publicPort":t[3]||(t[3]=e=>p.value=e),"onUpdate:isValidPublicPort":t[4]||(t[4]=e=>u.value=e),style:{width:"49%"}}),(0,l.Wm)(ye,{"onUpdate:selectedProtocol":t[5]||(t[5]=e=>r.value=e),style:{width:"49%"}})])])),_:1}),(0,l.Wm)(n,{class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(h,{bordered:""},{default:(0,l.w5)((()=>[(0,l.Wm)(m,{"expand-separator":"",icon:"settings",label:e.$t("mappingsCreateDialog.advancedSettingsLabel"),class:"mappings-create-dialog-icon"},{default:(0,l.w5)((()=>[(0,l.Wm)(s,{class:"mappings-create-dialog-advances-bg q-ma-md"},{default:(0,l.w5)((()=>[(0,l.Wm)(n,null,{default:(0,l.w5)((()=>[(0,l.Wm)(me,{"onUpdate:selectedAccount":t[6]||(t[6]=e=>c.value=e)})])),_:1})])),_:1})])),_:1},8,["label"])])),_:1})])),_:1}),(0,l.Wm)(f,{align:"between",class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(H.Z,{label:e.$t("mappingsCreateDialog.cancelBtn"),class:"q-mr-sm q-px-xl",color:"grey-8",onClick:t[7]||(t[7]=e=>b())},null,8,["label"]),(0,l.Wm)(H.Z,{disable:!1===v.value,color:"primary",icon:"add",label:e.$t("mappingsCreateDialog.createBtn"),onClick:t[8]||(t[8]=e=>w())},null,8,["disable","label"])])),_:1})])),_:1})])),_:1},8,["modelValue"])}}});var He=a(2074),Qe=a(3246),Le=a(1123),je=a(1821),ze=a(490);const Be=$e,Fe=Be;d()($e,"components",{QDialog:He.Z,QCard:y.Z,QBtn:c.Z,QCardSection:_.Z,QList:Qe.Z,QExpansionItem:Le.Z,QCardActions:je.Z,QItem:ze.Z});var Ue=a(9906),Ae=a(9302);const Ve=(0,l.aZ)({__name:"MappingsDeleteDialog",setup(e){const t=(0,ke.QT)().t,a=s(),n=(0,l.Fl)({get:()=>a.getShowDeleteMappingDialog,set:e=>{a.setShowDeleteMappingDialog(e)}}),i=(0,l.Fl)({get:()=>a.getKeyMappingsTable,set:e=>{a.setKeyMappingsTable(e)}}),o=(0,l.Fl)((()=>a.getSelectedMapping)),p=(0,l.Fl)((()=>{const e=(0,Ae.Z)();return e.dark.isActive?"/_/icons/bomb_dark.svg":"/_/icons/bomb_light.svg"}));function r(){n.value=!1}function c(){(0,We.Q)();const e=new Ce,a={mappingId:o.value.id};e.deleteMapping(a).then((()=>{(0,qe.LX)(`${t("mappingsDeleteDialog.deleteMappingSuccessfully")}`),i.value++,setTimeout((()=>{r()}),500)})).catch((e=>{console.error(e),(0,qe.s9)(e.response.data,`${t("mappingsDeleteDialog.deleteMappingError")}`)})).finally((()=>{(0,We.Z)()}))}return(e,t)=>((0,l.wg)(),(0,l.j4)(Ue.Z,{showDeleteDialog:n.value,"onUpdate:showDeleteDialog":t[2]||(t[2]=e=>n.value=e),titleDialog:e.$t("mappingsDeleteDialog.title"),imagePath:p.value,messageToDelete:e.$t("mappingsDeleteDialog.confirmRemoveMapping"),warningToDelete:e.$t("mappingsDeleteDialog.warningRemoveMapping")},{"card-actions":(0,l.w5)((()=>[(0,l.Wm)(H.Z,{label:e.$t("mappingsDeleteDialog.cancelBtn"),color:"grey-8",onClick:t[0]||(t[0]=e=>r())},null,8,["label"]),(0,l.Wm)(H.Z,{icon:"delete",label:e.$t("mappingsDeleteDialog.deleteBtn"),color:"negative",onClick:t[1]||(t[1]=e=>c())},null,8,["label"])])),_:1},8,["showDeleteDialog","titleDialog","imagePath","messageToDelete","warningToDelete"]))}}),Ye=Ve,Ke=Ye;var Re=a(4877),Ne=a(5699);const Ee={key:1},Xe={style:{"font-size":"0.9708rem"}},Ge={style:{"font-size":"1.1326rem"}},Je={class:"flex justify-start items-center"},Oe={style:{padding:"0px",margin:"0 4px"}},et={style:{"font-size":"1.1326rem"}},tt={style:{"padding-left":"8px"}},at={key:0,class:"q-my-sm"},lt=(0,l.aZ)({__name:"TargetContainerSelect",emits:["update:selectedContainer"],setup(e,{emit:t}){const a=t,o=(0,n.iH)(!1),s=(0,n.iH)(!1),p=(0,n.iH)([]),r=(0,n.iH)("");function c(){o.value=!0;const e=new Ne.Z;e.getContainers().then((e=>{p.value=[],0!==e.data.body.length&&(p.value=u(e.data.body))})).catch((e=>{console.error(e)})).finally((()=>{o.value=!1}))}function u(e){return e.map((e=>({label:e.hostname,value:e.id,portBindings:e.portBindings?e.portBindings:[]})))}return(0,l.bv)((()=>{c()})),(0,l.YP)(r,(()=>{!0!==s.value?a("update:selectedContainer",r.value):a("update:selectedContainer","")})),(e,t)=>{const a=(0,l.up)("q-skeleton"),n=(0,l.up)("q-icon");return!0===o.value?((0,l.wg)(),(0,l.j4)(a,{key:0,animation:"wave",height:"75px"})):((0,l.wg)(),(0,l.iD)("div",Ee,[(0,l.Wm)(Re.Z,{label:e.$t("targetContainerSelect.selectContainer"),selected:r.value,"onUpdate:selected":t[2]||(t[2]=e=>r.value=e),options:p.value},{"selected-content":(0,l.w5)((({selectedItem:e})=>[(0,l._)("div",Xe,(0,i.zw)(null===e||void 0===e?void 0:e.label),1)])),"option-content":(0,l.w5)((({option:a})=>[a.portBindings.length>0?((0,l.wg)(),(0,l.iD)("div",{key:0,onClick:t[0]||(t[0]=e=>s.value=!1)},[(0,l._)("div",Ge,(0,i.zw)(a.label),1),(0,l._)("div",Je,[((0,l.wg)(!0),(0,l.iD)(l.HY,null,(0,l.Ko)(a.portBindings,((e,t)=>((0,l.wg)(),(0,l.iD)("div",{key:t},[(0,l.Wm)(q,{labelLeft:e.serviceName?e.serviceName:"N/A",labelRight:`${e.publicPort}/${e.protocol}:${e.privatePort}`,fontSize:"0.809rem",class:"q-ma-xs"},null,8,["labelLeft","labelRight"])])))),128))])])):((0,l.wg)(),(0,l.iD)("div",{key:1,class:"disabled",onClick:t[1]||(t[1]=e=>s.value=!0)},[(0,l._)("div",Oe,[(0,l._)("div",et,(0,i.zw)(a.label),1),(0,l._)("div",tt,(0,i.zw)(e.$t("targetContainerSelect.noPortBindings")),1)])]))])),_:1},8,["label","selected","options"]),!0===s.value?((0,l.wg)(),(0,l.iD)("div",at,[(0,l.Wm)(n,{name:"error",size:"1rem",color:"negative"}),(0,l.Uk)(" "+(0,i.zw)(e.$t("targetContainerSelect.noPortBindingsAlert")),1)])):(0,l.kq)("",!0)]))}}}),nt=lt,it=nt;d()(lt,"components",{QSkeleton:ie.ZP,QIcon:B.Z});const ot=(0,l.aZ)({__name:"TargetMappingSelect",props:{selectedMappingId:{}},emits:["update:selectedMappingId"],setup(e,{emit:t}){const a=e,i=s(),o=t,p=(0,n.iH)([]),r=(0,n.iH)(a.selectedMappingId),c=(0,n.iH)(0),u=(0,l.Fl)((()=>i.getMappingsList));function g(){return 0===u.value.length?(r.value="",[]):u.value.map((e=>({label:`#${e.id} - ${e.hostname?e.hostname:"*"}`,value:e.id.toString()})))}return(0,l.bv)((()=>{p.value=g(),c.value++})),(0,l.YP)(r,(()=>{o("update:selectedMappingId",r.value)})),(0,l.YP)((()=>a.selectedMappingId),(()=>{r.value=a.selectedMappingId}),{immediate:!0}),(e,t)=>((0,l.wg)(),(0,l.j4)(re.Z,{selected:r.value,"onUpdate:selected":t[0]||(t[0]=e=>r.value=e),label:e.$t("targetMappingSelect.selectMapping"),options:p.value,dense:!1,isDisabled:0===p.value.length,key:c.value},null,8,["selected","label","options","isDisabled"]))}}),st=ot,pt=st,rt={class:"flex justify-between items-center"},ct={class:"title-dialog"},ut=(0,l.aZ)({__name:"TargetCreateDialog",setup(e){const t=(0,ke.QT)().t,a=s(),o=(0,n.iH)(""),p=(0,n.iH)(""),r=(0,l.Fl)({get:()=>a.getSelectedMapping,set:e=>a.setSelectedMapping(e)}),c=(0,l.Fl)({get:()=>a.getShowCreateTargetDialog,set:e=>a.setShowCreateTargetDialog(e)}),u=(0,l.Fl)({get:()=>a.getKeyMappingsTable,set:e=>a.setKeyMappingsTable(e)}),g=(0,l.Fl)((()=>""!==o.value&&""!==p.value));function d(){r.value={},c.value=!1}function m(){(0,We.Q)();const e=new Ce,a={mappingId:parseInt(p.value),containerId:o.value};e.createTarget(a).then((()=>{(0,qe.LX)(t("targetCreateDialog.createTargetSuccessfully")),u.value++,d()})).catch((e=>{console.error(e),(0,qe.s9)(e.response.data,t("targetCreateDialog.createTargetError"))})).finally((()=>{(0,We.Z)()}))}return(0,l.bv)((()=>{r.value.id&&(p.value=r.value.id.toString())})),(e,t)=>{const a=(0,l.up)("q-btn"),n=(0,l.up)("q-card-section"),s=(0,l.up)("q-card-actions"),r=(0,l.up)("q-card"),u=(0,l.up)("q-dialog");return(0,l.wg)(),(0,l.j4)(u,{modelValue:c.value,"onUpdate:modelValue":t[5]||(t[5]=e=>c.value=e),persistent:""},{default:(0,l.w5)((()=>[(0,l.Wm)(r,{style:{width:"700px","max-width":"80vw"},class:"dialog-card-bg"},{default:(0,l.w5)((()=>[(0,l._)("div",rt,[(0,l._)("div",ct,(0,i.zw)(e.$t("targetCreateDialog.title")),1),(0,l.Wm)(a,{flat:"",round:"",dense:"",icon:"close",onClick:t[0]||(t[0]=e=>d())})]),(0,l.Wm)(n,{class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(pt,{selectedMappingId:p.value,"onUpdate:selectedMappingId":t[1]||(t[1]=e=>p.value=e)},null,8,["selectedMappingId"])])),_:1}),(0,l.Wm)(n,{class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(it,{"onUpdate:selectedContainer":t[2]||(t[2]=e=>o.value=e)})])),_:1}),(0,l.Wm)(s,{align:"between",class:"q-px-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(H.Z,{label:e.$t("targetCreateDialog.cancelBtn"),class:"q-mr-sm q-px-xl",color:"grey-8",onClick:t[3]||(t[3]=e=>d())},null,8,["label"]),(0,l.Wm)(H.Z,{color:"primary",icon:"add",label:e.$t("targetCreateDialog.createBtn"),disable:!1===g.value,onClick:t[4]||(t[4]=e=>m())},null,8,["label","disable"])])),_:1})])),_:1})])),_:1},8,["modelValue"])}}}),gt=ut,dt=gt;d()(ut,"components",{QDialog:He.Z,QCard:y.Z,QBtn:c.Z,QCardSection:_.Z,QCardActions:je.Z,QItem:ze.Z});const mt=(0,l.aZ)({__name:"TargetDeleteDialog",setup(e){const t=(0,ke.QT)().t,a=s(),n=(0,l.Fl)({get:()=>a.getShowDeleteTargetDialog,set:e=>{a.setShowDeleteTargetDialog(e)}}),i=(0,l.Fl)({get:()=>a.getKeyMappingsTable,set:e=>{a.setKeyMappingsTable(e)}}),o=(0,l.Fl)((()=>a.getSelectedTarget)),p=(0,l.Fl)((()=>{const e=(0,Ae.Z)();return e.dark.isActive?"/_/icons/bomb_dark.svg":"/_/icons/bomb_light.svg"}));function r(){n.value=!1}function c(){(0,We.Q)();const e=new Ce,a={targetId:o.value.id,mappingId:o.value.mappingId};e.deleteTarget(a).then((()=>{(0,qe.LX)(`${t("targetDeleteDialog.deleteTargetSuccessfully")}`),i.value++,setTimeout((()=>{r()}),500)})).catch((e=>{console.error(e),(0,qe.s9)(e.response.data,`${t("targetDeleteDialog.deleteTargetError")}`)})).finally((()=>{(0,We.Z)()}))}return(e,t)=>((0,l.wg)(),(0,l.j4)(Ue.Z,{showDeleteDialog:n.value,"onUpdate:showDeleteDialog":t[2]||(t[2]=e=>n.value=e),titleDialog:e.$t("targetDeleteDialog.title"),imagePath:p.value,messageToDelete:e.$t("targetDeleteDialog.confirmRemoveTarget"),warningToDelete:e.$t("targetDeleteDialog.warningRemoveTarget")},{"card-actions":(0,l.w5)((()=>[(0,l.Wm)(H.Z,{label:e.$t("targetDeleteDialog.cancelBtn"),color:"grey-8",onClick:t[0]||(t[0]=e=>r())},null,8,["label"]),(0,l.Wm)(H.Z,{label:e.$t("targetDeleteDialog.deleteBtn"),color:"negative",onClick:t[1]||(t[1]=e=>c())},null,8,["label"])])),_:1},8,["showDeleteDialog","titleDialog","imagePath","messageToDelete","warningToDelete"]))}}),vt=mt,bt=vt;var wt=function(e,t,a,l){function n(e){return e instanceof a?e:new a((function(t){t(e)}))}return new(a||(a=Promise))((function(a,i){function o(e){try{p(l.next(e))}catch(t){i(t)}}function s(e){try{p(l["throw"](e))}catch(t){i(t)}}function p(e){e.done?a(e.value):n(e.value).then(o,s)}p((l=l.apply(e,t||[])).next())}))};const ht=(0,l.aZ)({__name:"MappingsIndex",setup(e){const t=s(),a=(0,n.iH)([]),i=(0,n.iH)(0),o=(0,n.iH)(0),p=(0,n.iH)(!1),r=(0,l.Fl)({get:()=>t.getMappingsList,set:e=>t.setMappingsList(e)}),c=(0,l.Fl)((()=>t.getShowCreateMappingDialog)),u=(0,l.Fl)((()=>t.getShowCreateTargetDialog)),g=(0,l.Fl)((()=>t.getKeyMappingsTable));function d(){return wt(this,void 0,void 0,(function*(){try{p.value=!0;const e=new Ce,t=yield e.getMappings();if(0===t.data.body.length)return;r.value=t.data.body;const l=yield m();a.value=v(r.value,l)}catch(e){console.error(e)}finally{p.value=!1}}))}function m(){return wt(this,void 0,void 0,(function*(){try{const e=new ce.Z,t=yield e.getAccounts();return 0===t.data.body.length?[]:t.data.body}catch(e){return console.error(e),[]}}))}function v(e,t){return e.map((e=>{const a=t.find((t=>t.id===e.accountId));return{id:e.id,accountId:e.accountId,accountName:(null===a||void 0===a?void 0:a.username)||"",hostname:e.hostname,protocol:e.protocol,publicPort:e.publicPort,targets:e.targets}}))}return(0,l.YP)(c,(()=>{i.value++})),(0,l.YP)(u,(()=>{o.value++})),(0,l.YP)(g,(()=>{a.value=[],r.value=[],d()})),(0,l.bv)((()=>{d()})),(e,t)=>{const n=(0,l.up)("q-page");return(0,l.wg)(),(0,l.j4)(n,{padding:""},{default:(0,l.w5)((()=>[(0,l.Wm)(Ke),((0,l.wg)(),(0,l.j4)(Fe,{key:i.value})),(0,l.Wm)(bt),((0,l.wg)(),(0,l.j4)(dt,{key:o.value})),!0===p.value?((0,l.wg)(),(0,l.j4)(pe,{key:0})):((0,l.wg)(),(0,l.j4)(K,{key:1,mappingsList:a.value},null,8,["mappingsList"]))])),_:1})}}});var ft=a(9885);const Dt=ht,yt=Dt;d()(ht,"components",{QPage:ft.Z})}}]);