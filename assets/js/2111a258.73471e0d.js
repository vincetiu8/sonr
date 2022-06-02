"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[5186],{3905:function(e,t,r){r.d(t,{Zo:function(){return l},kt:function(){return h}});var n=r(7294);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,i=function(e,t){if(null==e)return{};var r,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var c=n.createContext({}),d=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},l=function(e){var t=d(e.components);return n.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var r=e.components,i=e.mdxType,o=e.originalType,c=e.parentName,l=s(e,["components","mdxType","originalType","parentName"]),u=d(r),h=i,f=u["".concat(c,".").concat(h)]||u[h]||p[h]||o;return r?n.createElement(f,a(a({ref:t},l),{},{components:r})):n.createElement(f,a({ref:t},l))}));function h(e,t){var r=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=r.length,a=new Array(o);a[0]=u;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var d=2;d<o;d++)a[d]=r[d];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}u.displayName="MDXCreateElement"},393:function(e,t,r){r.r(t),r.d(t,{assets:function(){return l},contentTitle:function(){return c},default:function(){return h},frontMatter:function(){return s},metadata:function(){return d},toc:function(){return p}});var n=r(7462),i=r(3366),o=(r(7294),r(3905)),a=["components"],s={title:"Objects",id:"objects",displayed_sidebar:"highwaySidebar"},c=void 0,d={unversionedId:"highway-sdk/objects",id:"highway-sdk/objects",title:"Objects",description:"The Sonr object module is used to store the records of verifiable objects for a specific application powered by the Sonr Network. Each record contains an ObjectDoc which describes the type definition of the associated object.",source:"@site/articles/highway-sdk/objects.md",sourceDirName:"highway-sdk",slug:"/highway-sdk/objects",permalink:"/articles/highway-sdk/objects",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/highway-sdk/objects.md",tags:[],version:"current",frontMatter:{title:"Objects",id:"objects",displayed_sidebar:"highwaySidebar"},sidebar:"highwaySidebar",previous:{title:"Registry",permalink:"/articles/highway-sdk/registry"},next:{title:"Channels",permalink:"/articles/highway-sdk/channels"}},l={},p=[{value:"Overview",id:"overview",level:2},{value:"Usage",id:"usage",level:2},{value:"<code>CreateObject()</code> - Register&#39;s a new type definition for a given application",id:"createobject---registers-a-new-type-definition-for-a-given-application",level:3},{value:"<code>ReadObject()</code> - Returns the WhatIs record for a provided object",id:"readobject---returns-the-whatis-record-for-a-provided-object",level:3},{value:"<code>UpdateObject()</code> - Changes the configuration of the verifiable object",id:"updateobject---changes-the-configuration-of-the-verifiable-object",level:3},{value:"Status Codes",id:"status-codes",level:2}],u={toc:p};function h(e){var t=e.components,r=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("p",null,"#Objects\nThe Sonr object module is used to store the records of verifiable objects for a specific application powered by the Sonr Network. Each record contains an ",(0,o.kt)("inlineCode",{parentName:"p"},"ObjectDoc")," which describes the type definition of the associated object."),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"The record type utilized in the ",(0,o.kt)("strong",{parentName:"p"},"Object module")," is the ",(0,o.kt)("inlineCode",{parentName:"p"},"WhatIs")," type. This type provides both an interface to utilize VerifiableCredentials and modify the ObjectDoc type definition"),(0,o.kt)("h2",{id:"usage"},"Usage"),(0,o.kt)("blockquote",null,(0,o.kt)("p",{parentName:"blockquote"},"Blockchain Methods supplied by Object Module. Full implementation is still a work in progress.")),(0,o.kt)("h3",{id:"createobject---registers-a-new-type-definition-for-a-given-application"},(0,o.kt)("inlineCode",{parentName:"h3"},"CreateObject()")," - Register's a new type definition for a given application"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-Text"},"- (`string`) Creator            : The Account Address signing this message\n- (`Session`) Session           : The Session for the authenticated user\n- (`string`) Label              : Name of the Object defined by developer\n- (`string`) Description        : Description of the object defined by developer\n- (`List`) InitialFields        : List of the initial `ObjectField`s to create\n")),(0,o.kt)("h3",{id:"readobject---returns-the-whatis-record-for-a-provided-object"},(0,o.kt)("inlineCode",{parentName:"h3"},"ReadObject()")," - Returns the WhatIs record for a provided object"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-Text"},"- (`string`) Creator                : The Account Address signing this message\n- (`Session`) Session               : The Session for the authenticated user\n- (`string`) Did                    : The DID of the object to read\n")),(0,o.kt)("h3",{id:"updateobject---changes-the-configuration-of-the-verifiable-object"},(0,o.kt)("inlineCode",{parentName:"h3"},"UpdateObject()")," - Changes the configuration of the verifiable object"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-Text"},"- (`string`) Creator            : The Account Address signing this message\n- (`Session`) Session           : The Session for the authenticated user\n- (`string`) Label              : Name of the Object defined by developer\n- (`string`) Description        : Description of the object defined by developer\n- (`List`) AddedFields          : List of the `ObjectField`s to add\n- (`List`) RemovedFields        : List of the `ObjectField`s to remove\n")),(0,o.kt)("h2",{id:"status-codes"},"Status Codes"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"200 - SUCCESS\n300 - MULTIPLE CHOICE\n304 - NOT MODIFIED\n400 - BAD REQUEST\n401 - NOT AUTHORIZED\n\n")))}h.isMDXComponent=!0}}]);