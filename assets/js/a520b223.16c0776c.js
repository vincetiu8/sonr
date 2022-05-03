"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[8348],{3905:function(e,n,t){t.d(n,{Zo:function(){return u},kt:function(){return f}});var r=t(7294);function o(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function a(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){o(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function c(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var s=r.createContext({}),d=function(e){var n=r.useContext(s),t=n;return e&&(t="function"==typeof e?e(n):a(a({},n),e)),t},u=function(e){var n=d(e.components);return r.createElement(s.Provider,{value:n},e.children)},l={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},p=r.forwardRef((function(e,n){var t=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),p=d(t),f=o,b=p["".concat(s,".").concat(f)]||p[f]||l[f]||i;return t?r.createElement(b,a(a({ref:n},u),{},{components:t})):r.createElement(b,a({ref:n},u))}));function f(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var i=t.length,a=new Array(i);a[0]=p;var c={};for(var s in n)hasOwnProperty.call(n,s)&&(c[s]=n[s]);c.originalType=e,c.mdxType="string"==typeof e?e:o,a[1]=c;for(var d=2;d<i;d++)a[d]=t[d];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}p.displayName="MDXCreateElement"},8601:function(e,n,t){t.r(n),t.d(n,{assets:function(){return u},contentTitle:function(){return s},default:function(){return f},frontMatter:function(){return c},metadata:function(){return d},toc:function(){return l}});var r=t(7462),o=t(3366),i=(t(7294),t(3905)),a=["components"],c={title:"Access & Authentication",id:"access-authentication",displayed_sidebar:"buildSidebar"},s=void 0,d={unversionedId:"motor-node/access-authentication",id:"motor-node/access-authentication",title:"Access & Authentication",description:"We think authentication should be simple, yet secure. The Sonr network uses s.",source:"@site/articles/motor-node/access-authentication.md",sourceDirName:"motor-node",slug:"/motor-node/access-authentication",permalink:"/sonr/articles/motor-node/access-authentication",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/motor-node/access-authentication.md",tags:[],version:"current",frontMatter:{title:"Access & Authentication",id:"access-authentication",displayed_sidebar:"buildSidebar"},sidebar:"buildSidebar",previous:{title:"The Sonr Stack",permalink:"/sonr/articles/sonr-stack"},next:{title:"Discovery",permalink:"/sonr/articles/motor-node/discovery"}},u={},l=[],p={toc:l};function f(e){var n=e.components,t=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,r.Z)({},p,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"We think authentication should be simple, yet secure. The Sonr network uses s."),(0,i.kt)("h1",{id:"platform-credentials"},"Platform Credentials"),(0,i.kt)("p",null,"Currently, our implementations of Webauthn use the 'platform-specific' credential options meaning our servers will request your operating system to use whichever authentication method is most native to it. For information on what authentication mechanisims are supported"),(0,i.kt)("h1",{id:"did-registration"},"DID Registration"),(0,i.kt)("p",null,"When a user registers their domain will be prompted to supply your user credentials for relating to our generated ",(0,i.kt)("inlineCode",{parentName:"p"},"DID")," which is then paired with provided ",(0,i.kt)("inlineCode",{parentName:"p"},"PublicKeyCredentials")),(0,i.kt)("p",null,"The folowing is an example of a generated ",(0,i.kt)("inlineCode",{parentName:"p"},"WhoIs")," which repersented a user in our registry"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-javascript"},'{\n  "@context": [\n    "https://www.w3.org/ns/did/v1"\n  ],\n  "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",\n  "controller": [\n    "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",\n    "did:sonr:f03a00f1-9615-4060-bd00-bd282e150c46"\n  ],\n  "verificationMethod": [\n    {\n      "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1",\n      "type": "JsonWebKey2020",\n      "controller": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",\n      "publicKeyJwk": {\n        "kty": "EC",\n        "crv": "P-256",\n        "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",\n        "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"\n      }\n    },\n    {\n      "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-2",\n      "type": "JsonWebKey2020",\n      "controller": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",\n      "publicKeyJwk": {\n        "kty": "EC",\n        "crv": "P-256",\n        "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",\n        "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"\n      }\n    },\n    {\n      "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#added-assertion-method-1",\n      "controller": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",\n      "publicKeyBase58": "GGRj8PAR5tRgD5xqAhPna1bLa3UoYuxNEEhRmcYCPBm5",\n      "type": "Ed25519VerificationKey2018"\n    }\n  ],\n  "authentication": [\n    "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1",\n    "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-2"\n  ],\n  "assertionMethod": [\n    "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1"\n  ],\n  "service": [\n    {\n      "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-1",\n      "type": "sonr:bolt:eoverdracht",\n      "serviceEndpoint": "did:sonr:<vendor>#service-76"\n    },\n    {\n      "id": "did:sonr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-2",\n      "type": "sonr:core:consent",\n      "serviceEndpoint": "did:sonr:<vendor>#service-2"\n    }\n  ]\n}\n')),(0,i.kt)("p",null,"``"))}f.isMDXComponent=!0}}]);