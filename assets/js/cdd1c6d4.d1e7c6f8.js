"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[4436],{3905:function(e,n,t){t.d(n,{Zo:function(){return c},kt:function(){return w}});var r=t(7294);function o(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function s(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach((function(n){o(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function i(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var l=r.createContext({}),p=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):s(s({},n),e)),t},c=function(e){var n=p(e.components);return r.createElement(l.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},f=r.forwardRef((function(e,n){var t=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),f=p(t),w=o,d=f["".concat(l,".").concat(w)]||f[w]||u[w]||a;return t?r.createElement(d,s(s({ref:n},c),{},{components:t})):r.createElement(d,s({ref:n},c))}));function w(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var a=t.length,s=new Array(a);s[0]=f;var i={};for(var l in n)hasOwnProperty.call(n,l)&&(i[l]=n[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,s[1]=i;for(var p=2;p<a;p++)s[p]=t[p];return r.createElement.apply(null,s)}return r.createElement.apply(null,t)}f.displayName="MDXCreateElement"},9314:function(e,n,t){t.r(n),t.d(n,{assets:function(){return c},contentTitle:function(){return l},default:function(){return w},frontMatter:function(){return i},metadata:function(){return p},toc:function(){return u}});var r=t(7462),o=t(3366),a=(t(7294),t(3905)),s=["components"],i={title:"HowIs",slug:"3M4V-howis",createdAt:new Date("2022-04-26T14:53:09.000Z"),updatedAt:new Date("2022-04-28T19:49:44.000Z")},l="HowIs",p={unversionedId:"reference/highway-api/howis",id:"reference/highway-api/howis",title:"HowIs",description:"Overview",source:"@site/articles/reference/highway-api/howis.md",sourceDirName:"reference/highway-api",slug:"/reference/highway-api/3M4V-howis",permalink:"/articles/reference/highway-api/3M4V-howis",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/reference/highway-api/howis.md",tags:[],version:"current",frontMatter:{title:"HowIs",slug:"3M4V-howis",createdAt:"2022-04-26T14:53:09.000Z",updatedAt:"2022-04-28T19:49:44.000Z"}},c={},u=[{value:"Overview",id:"overview",level:2},{value:"Usage",id:"usage",level:2},{value:"Create HowIs",id:"create-howis",level:3},{value:"Response from CreateHowIs",id:"response-from-createhowis",level:3},{value:"Update HowIs",id:"update-howis",level:3},{value:"Response from UpdateHowIs",id:"response-from-updatehowis",level:3},{value:"Deleting HowIs",id:"deleting-howis",level:3},{value:"Response from DeleteHowIs",id:"response-from-deletehowis",level:3}],f={toc:u};function w(e){var n=e.components,t=(0,o.Z)(e,s);return(0,a.kt)("wrapper",(0,r.Z)({},f,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"howis"},"HowIs"),(0,a.kt)("h2",{id:"overview"},"Overview"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"HowIs"),"Object function is included within ",(0,a.kt)("inlineCode",{parentName:"p"},"Channel")," types. The ",(0,a.kt)("inlineCode",{parentName:"p"},"HowIs")," Object illustrates how a channel is associated with a ",(0,a.kt)("inlineCode",{parentName:"p"},"registry")," entry on the Sonr Network. A ",(0,a.kt)("inlineCode",{parentName:"p"},"registry")," entry on the Sonr Network can either be for an ",(0,a.kt)("inlineCode",{parentName:"p"},"application")," or a ",(0,a.kt)("inlineCode",{parentName:"p"},"name")),(0,a.kt)("h2",{id:"usage"},"Usage"),(0,a.kt)("h3",{id:"create-howis"},"Create HowIs"),(0,a.kt)("p",null,"The following is an example of a request to make a new HowIs"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgCreateHowIs {\n  string creator;\n  string did;\n  ChannelDoc channel;\n}\n")),(0,a.kt)("h3",{id:"response-from-createhowis"},"Response from CreateHowIs"),(0,a.kt)("p",null,"The following is an example of a response after creating a new HowIs"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgCreateHowIsResponse {\n    int32 code      // Code of the response\n    string message  // Message of the response\n    HowIs how_is    // HowIs of the Channel\n}\n")),(0,a.kt)("h3",{id:"update-howis"},"Update HowIs"),(0,a.kt)("p",null,"The following is an example of a request to update a HowIs"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgUpdateHowIs {\n  string creator;\n  string did;\n  ChannelDoc channel;\n}\n")),(0,a.kt)("h3",{id:"response-from-updatehowis"},"Response from UpdateHowIs"),(0,a.kt)("p",null,"The following is an example of a response after Updating a HowIs"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgUpdateHowIsResponse {\n    int32 code     // Code of the response\n    string message // Message of the response\n    HowIs how_is   // HowIs of the Channel\n}\n")),(0,a.kt)("h3",{id:"deleting-howis"},"Deleting HowIs"),(0,a.kt)("p",null,"The following is an example of a request to Delete a HowIs"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgDeleteHowIs {\n  string creator;\n  string did;\n}\n")),(0,a.kt)("h3",{id:"response-from-deletehowis"},"Response from DeleteHowIs"),(0,a.kt)("p",null,"The following is an example of a response after Deleting a How Is"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"MsgDeleteHowIsResponse {\n    int32 code     // Code of the response\n    string message // Message of the response\n}\n")))}w.isMDXComponent=!0}}]);