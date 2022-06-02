"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[2572],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return d}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function a(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var c=n.createContext({}),u=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},p=function(e){var t=u(e.components);return n.createElement(c.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),f=u(r),d=o,h=f["".concat(c,".").concat(d)]||f[d]||l[d]||i;return r?n.createElement(h,s(s({ref:t},p),{},{components:r})):n.createElement(h,s({ref:t},p))}));function d(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,s=new Array(i);s[0]=f;var a={};for(var c in t)hasOwnProperty.call(t,c)&&(a[c]=t[c]);a.originalType=e,a.mdxType="string"==typeof e?e:o,s[1]=a;for(var u=2;u<i;u++)s[u]=r[u];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},1473:function(e,t,r){r.r(t),r.d(t,{assets:function(){return p},contentTitle:function(){return c},default:function(){return d},frontMatter:function(){return a},metadata:function(){return u},toc:function(){return l}});var n=r(7462),o=r(3366),i=(r(7294),r(3905)),s=["components"],a={title:"Session",slug:"dsi2-session",createdAt:new Date("2022-04-26T14:49:33.000Z"),updatedAt:new Date("2022-04-27T21:55:57.000Z")},c=void 0,u={unversionedId:"reference/highway-api/session",id:"reference/highway-api/session",title:"Session",description:"Overview",source:"@site/articles/reference/highway-api/session.md",sourceDirName:"reference/highway-api",slug:"/reference/highway-api/dsi2-session",permalink:"/articles/reference/highway-api/dsi2-session",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/reference/highway-api/session.md",tags:[],version:"current",frontMatter:{title:"Session",slug:"dsi2-session",createdAt:"2022-04-26T14:49:33.000Z",updatedAt:"2022-04-27T21:55:57.000Z"}},p={},l=[],f={toc:l};function d(e){var t=e.components,r=(0,o.Z)(e,s);return(0,i.kt)("wrapper",(0,n.Z)({},f,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"#Sessions\n##Overview\n",(0,i.kt)("inlineCode",{parentName:"p"},"Sessions")," are used to approve a user's authenication status with each request made. For a request to be processed it must be associated with a valid session in relation to the channel. In essence, the Session makes sure that a user is approved to further interact with the Sonr Network. For more information on sessions see ",(0,i.kt)("inlineCode",{parentName:"p"},"Registry")),(0,i.kt)("p",null,"##Usage"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"Session {\n    string base_did;        // Base DID is the current Account or Application whois DID url\n    WhoIs whois;            // WhoIs is the current Document for the DID\n    Credential credential;  // Credential is the current Credential for the DID\n}\n")))}d.isMDXComponent=!0}}]);