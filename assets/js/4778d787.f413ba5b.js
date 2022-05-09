"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[643],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return d}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function a(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var c=n.createContext({}),u=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=u(e.components);return n.createElement(c.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,s=e.originalType,c=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),f=u(r),d=o,h=f["".concat(c,".").concat(d)]||f[d]||l[d]||s;return r?n.createElement(h,i(i({ref:t},p),{},{components:r})):n.createElement(h,i({ref:t},p))}));function d(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var s=r.length,i=new Array(s);i[0]=f;var a={};for(var c in t)hasOwnProperty.call(t,c)&&(a[c]=t[c]);a.originalType=e,a.mdxType="string"==typeof e?e:o,i[1]=a;for(var u=2;u<s;u++)i[u]=r[u];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},8121:function(e,t,r){r.r(t),r.d(t,{assets:function(){return p},contentTitle:function(){return c},default:function(){return d},frontMatter:function(){return a},metadata:function(){return u},toc:function(){return l}});var n=r(7462),o=r(3366),s=(r(7294),r(3905)),i=["components"],a={title:"Sessions",slug:"GYaU-sessions",createdAt:new Date("2022-04-27T18:21:38.000Z"),updatedAt:new Date("2022-04-27T18:21:43.000Z")},c=void 0,u={unversionedId:"reference/highway-api/sessions",id:"reference/highway-api/sessions",title:"Sessions",description:"Overview",source:"@site/articles/reference/highway-api/sessions.md",sourceDirName:"reference/highway-api",slug:"/reference/highway-api/GYaU-sessions",permalink:"/articles/reference/highway-api/GYaU-sessions",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/reference/highway-api/sessions.md",tags:[],version:"current",frontMatter:{title:"Sessions",slug:"GYaU-sessions",createdAt:"2022-04-27T18:21:38.000Z",updatedAt:"2022-04-27T18:21:43.000Z"}},p={},l=[],f={toc:l};function d(e){var t=e.components,r=(0,o.Z)(e,i);return(0,s.kt)("wrapper",(0,n.Z)({},f,r,{components:t,mdxType:"MDXLayout"}),(0,s.kt)("p",null,"#Sessions\n##Overview\nA ",(0,s.kt)("inlineCode",{parentName:"p"},"Session")," is used to denote a user's authenication status with each request made. In order for a request to be processed it must be associated with a valid session in relation to the channel.  In essence, the Session makes sure that a user is valid to further interact with the Sonr Network. For more information on sessions see ",(0,s.kt)("inlineCode",{parentName:"p"},"Registry")),(0,s.kt)("p",null,"##Usage"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"Session {\n    string base_did = 1;        // Base DID is the current Account or Application whois DID url\n    WhoIs whois = 2;            // WhoIs is the current Document for the DID\n    Credential credential = 3;  // Credential is the current Credential for the DID\n}\n")))}d.isMDXComponent=!0}}]);