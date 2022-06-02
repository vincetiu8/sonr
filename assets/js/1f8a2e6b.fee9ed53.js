"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[6181],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return h}});var o=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function r(e,t){if(null==e)return{};var n,o,a=function(e,t){if(null==e)return{};var n,o,a={},l=Object.keys(e);for(o=0;o<l.length;o++)n=l[o],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(o=0;o<l.length;o++)n=l[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=o.createContext({}),c=function(e){var t=o.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=c(e.components);return o.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},p=o.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,u=r(e,["components","mdxType","originalType","parentName"]),p=c(n),h=a,k=p["".concat(s,".").concat(h)]||p[h]||d[h]||l;return n?o.createElement(k,i(i({ref:t},u),{},{components:n})):o.createElement(k,i({ref:t},u))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,i=new Array(l);i[0]=p;var r={};for(var s in t)hasOwnProperty.call(t,s)&&(r[s]=t[s]);r.originalType=e,r.mdxType="string"==typeof e?e:a,i[1]=r;for(var c=2;c<l;c++)i[c]=n[c];return o.createElement.apply(null,i)}return o.createElement.apply(null,n)}p.displayName="MDXCreateElement"},5328:function(e,t,n){n.r(t),n.d(t,{assets:function(){return u},contentTitle:function(){return s},default:function(){return h},frontMatter:function(){return r},metadata:function(){return c},toc:function(){return d}});var o=n(7462),a=n(3366),l=(n(7294),n(3905)),i=["components"],r={title:"ADR-005 - WIP",id:"adr-005",displayed_sidebar:"resourcesSidebar"},s="ADR-005 Non-Fungible Token Standard for Objects (WIP)",c={unversionedId:"reference/adr-005",id:"reference/adr-005",title:"ADR-005 - WIP",description:"Core Values of ADR-005",source:"@site/articles/reference/ADR-005.md",sourceDirName:"reference",slug:"/reference/adr-005",permalink:"/articles/reference/adr-005",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/articles/reference/ADR-005.md",tags:[],version:"current",frontMatter:{title:"ADR-005 - WIP",id:"adr-005",displayed_sidebar:"resourcesSidebar"},sidebar:"resourcesSidebar",previous:{title:"ADR-004 - WIP",permalink:"/articles/reference/adr-004"},next:{title:"ADR-006 - WIP",permalink:"/articles/reference/adr-006"}},u={},d=[{value:"1. What is a Non-Fungible Token",id:"1-what-is-a-non-fungible-token",level:2},{value:"1.1 How Non-Fungible Tokens will Work on the Sonr Ecosystem",id:"11-how-non-fungible-tokens-will-work-on-the-sonr-ecosystem",level:3},{value:"1.2 How Non-Fungible Tokens will be utilized to create a consumer focused Market",id:"12-how-non-fungible-tokens-will-be-utilized-to-create-a-consumer-focused-market",level:3},{value:"1.3 How Non-Fungible Tokens can interact with the ecosystem",id:"13-how-non-fungible-tokens-can-interact-with-the-ecosystem",level:3},{value:"2. How will Non-Fungible Tokens be purchased, sold &amp; transferred on the Sonr blockchain",id:"2-how-will-non-fungible-tokens-be-purchased-sold--transferred-on-the-sonr-blockchain",level:2},{value:"2.1 Base of the CW721 Specification",id:"21-base-of-the-cw721-specification",level:3},{value:"2.2 Metadata with the Non-Fungible Smart Contracts (CW721)",id:"22-metadata-with-the-non-fungible-smart-contracts-cw721",level:3},{value:"2.3 How Gas Fees will allow the Blockchain &amp; Validators to recieve more token",id:"23-how-gas-fees-will-allow-the-blockchain--validators-to-recieve-more-token",level:3},{value:"2.4 How Non-Fungible Tokens can be Purchased, Sold, &amp; Transferred",id:"24-how-non-fungible-tokens-can-be-purchased-sold--transferred",level:3},{value:"3. Who will be using, creating and purchasing Non-Fungible Tokens",id:"3-who-will-be-using-creating-and-purchasing-non-fungible-tokens",level:2},{value:"4. Specification for Non-Fungible Tokens on the Blockchain",id:"4-specification-for-non-fungible-tokens-on-the-blockchain",level:2},{value:"Goals of Implementation are as follows",id:"goals-of-implementation-are-as-follows",level:3},{value:"x/nft Module Required Functions",id:"xnft-module-required-functions",level:3},{value:"Types for x/nft Module",id:"types-for-xnft-module",level:3},{value:"5. CW721 Specification In Detail",id:"5-cw721-specification-in-detail",level:2},{value:"CW721-base Specfication",id:"cw721-base-specfication",level:3},{value:"Base Messages",id:"base-messages",level:2},{value:"Base Queries",id:"base-queries",level:3},{value:"Base Receiver",id:"base-receiver",level:3},{value:"Metadata Queries",id:"metadata-queries",level:3}],p={toc:d};function h(e){var t=e.components,n=(0,a.Z)(e,i);return(0,l.kt)("wrapper",(0,o.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"adr-005-non-fungible-token-standard-for-objects-wip"},"ADR-005 Non-Fungible Token Standard for Objects (WIP)"),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Core Values of ADR-005")),(0,l.kt)("ol",null,(0,l.kt)("li",{parentName:"ol"},"What is a Non-Fungible Token\n1.1 How Non-Fungible Tokens will Work on the Sonr Ecosystem\n1.2 How Non-Fungible Tokens will be utilized to create a Consumer Focused Market\n1.3 How Non-Fungible Tokens can Interact with the Ecosystem"),(0,l.kt)("li",{parentName:"ol"},"How will Non-Fungible Tokens can be purchased, sold & transferred on the Sonr blockchain\n2.1 Base of the CW721 Specification\n2.2 Metadata with the Non-Fungible Smart Contracts (CW721)\n2.3 How Gas Fees will allow the Blockchain to recieve more token\n2.4 How Non-Fungible Tokens can be Purchased, Sold & Transferred"),(0,l.kt)("li",{parentName:"ol"},"Who will be using, creating and purchasing Non-Fungible Tokens"),(0,l.kt)("li",{parentName:"ol"},"Specification for Non-Fungible Tokens on the Blockchain\n4.1 Goals of Implementation\n4.2 x/nft Module Required Functions\n4.3 Types for x/nft module"),(0,l.kt)("li",{parentName:"ol"},"CW721 Specification In Detail\n5.1 CW721 Base Spec\n5.2 Base Messages\n5.3 Base Queries\n5.4 Base Receiver\n5.5 Metadata Queries")),(0,l.kt)("h2",{id:"1-what-is-a-non-fungible-token"},"1. What is a Non-Fungible Token"),(0,l.kt)("p",null,"A Non-Fungible Token or (NFT) is a digital asset that usually pertains to art, in-game items, music videos, and tickets although it is not limited to these specific items listed above. NFT's can be looked at as a verifiable purchase of the items listed above."),(0,l.kt)("h3",{id:"11-how-non-fungible-tokens-will-work-on-the-sonr-ecosystem"},"1.1 How Non-Fungible Tokens will Work on the Sonr Ecosystem"),(0,l.kt)("p",null,"The Non-Fungible Tokens (NFT) will be owned, trasnferred and sold on a basic level. For example, let's say that User-A purchases an NFT that is a picture of User-B's cat. User-A now owns that specific NFT photo of User-B's cat and can hold it, sell it to another user (via marketplace), or transfer the NFT to another user."),(0,l.kt)("h3",{id:"12-how-non-fungible-tokens-will-be-utilized-to-create-a-consumer-focused-market"},"1.2 How Non-Fungible Tokens will be utilized to create a consumer focused Market"),(0,l.kt)("p",null,"NFT's will allow users of the Sonr ecosystem to buy and sell art, music, tickets, in-game items, and videos. These specific examples will allow B2B and B2C projects to come onto the market in a cost effective way due to Sonr's low gas fees."),(0,l.kt)("h3",{id:"13-how-non-fungible-tokens-can-interact-with-the-ecosystem"},"1.3 How Non-Fungible Tokens can interact with the ecosystem"),(0,l.kt)("p",null,"Non-Fungible tokens will allow the ecosystem to obtain fees via the purchasing, selling, and transferring of the Non-Fungible Tokens."),(0,l.kt)("h2",{id:"2-how-will-non-fungible-tokens-be-purchased-sold--transferred-on-the-sonr-blockchain"},"2. How will Non-Fungible Tokens be purchased, sold & transferred on the Sonr blockchain"),(0,l.kt)("p",null,"The Sonr Blockchain will accept smart contracts that are in the Cosmos WebAssembly (CosmWasm) which is a smart contracting platform built for the Cosmos Ecosystem. The CW721 Specification will allow developers to utilize Sonr's specific NFT implementation and have a specific set of rules to follow and allow developers to add custom logic."),(0,l.kt)("h3",{id:"21-base-of-the-cw721-specification"},"2.1 Base of the CW721 Specification"),(0,l.kt)("p",null,"The Base will handle ownership, transfers, and allowances. ",(0,l.kt)("em",{parentName:"p"},"Allowances pertain to users allowing a portion of their wallet to be spent to purchase NFT's as well as sell and transfer them.")),(0,l.kt)("h3",{id:"22-metadata-with-the-non-fungible-smart-contracts-cw721"},"2.2 Metadata with the Non-Fungible Smart Contracts (CW721)"),(0,l.kt)("p",null,"Metadata is used to give Non-Fungible Tokens the ability to show specific data on a blockchain. For Example, User-A's cat NFT's would all have metadata with an on-chain or off-chain metadata with the picture of the cat."),(0,l.kt)("h3",{id:"23-how-gas-fees-will-allow-the-blockchain--validators-to-recieve-more-token"},"2.3 How Gas Fees will allow the Blockchain & Validators to recieve more token"),(0,l.kt)("p",null,"Along with the purchase, sell or transfer of an NFT is an associated gas fee. These gas fees can be thought of as a reward for the computationally difficult process of verifying the purchase, sale or transfer of an NFT. These Gas Fees will be utilized to give the validators on the network encentive to stay on the network as validators."),(0,l.kt)("h3",{id:"24-how-non-fungible-tokens-can-be-purchased-sold--transferred"},"2.4 How Non-Fungible Tokens can be Purchased, Sold, & Transferred"),(0,l.kt)("p",null,"Purchasing NFT\u2019s will be on an NFT marketplace. These marketplaces can be made by any developer on the network and doesn\u2019t necessarily need to be owned by Sonr itself. Although Sonr may not own the majority (or any) of the marketplaces, Gas Fees spoken about above will still go to the Sonr blockchain and its validators."),(0,l.kt)("h2",{id:"3-who-will-be-using-creating-and-purchasing-non-fungible-tokens"},"3. Who will be using, creating and purchasing Non-Fungible Tokens"),(0,l.kt)("p",null,"Sonr users from all walks of life will be utilizing Non-Fungible tokens for various reasons. Developers will be able to take the opportunity to create elaborate smart contracts for Profile Picture (PfP) Projects, online games, music, videos, and more. Users will be able to take advantage of these projects such as PfP projects, play online games, purchase music, videos, and more. Artists and individuals will be empowered to share their art in various forms listed above via Non-Fungible Tokens with the world via the Sonr Blockchain."),(0,l.kt)("h2",{id:"4-specification-for-non-fungible-tokens-on-the-blockchain"},"4. Specification for Non-Fungible Tokens on the Blockchain"),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Implementation of Non-Fungible Tokens into Blockchain")),(0,l.kt)("p",null,"The implementation of the NFT module into our blockchain will provide the ability to create smart contracts based on the CW-721 specification. This specification will allow developers to create smart contracts based on types that can be called upon by developers with their own set of logic. ",(0,l.kt)("em",{parentName:"p"},"Not all of the functionality may be used in the CW-721 specification but is available to be used by all developers.")),(0,l.kt)("h3",{id:"goals-of-implementation-are-as-follows"},"Goals of Implementation are as follows"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"Store Non-Fungible Tokens and track their ownership via the blockchain"),(0,l.kt)("li",{parentName:"ul"},"Expose Keeper interface for composing modules to trasnfer, mint, and burn Non-Fungible tokens."),(0,l.kt)("li",{parentName:"ul"},"Expose external Message interface for users to purchase, sell and transfer ownership of their Non-Fungible Tokens."),(0,l.kt)("li",{parentName:"ul"},"Query Non-Fungible Tokens for their supply information, ownership of a given NFT, and a list of operators that can access a specific owners tokens.")),(0,l.kt)("h3",{id:"xnft-module-required-functions"},"x/nft Module Required Functions"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"MsgNewClass")," - Receive the user's request to create a class, and call the NewClass of the x/nft module."),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"MsgUpdateClass")," - Receive the user's request to update a class, and call the UpdateClass of the x/nft module."),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"MsgMintNFT")," - Rceive the user's request to mint an NFT, and call the MintNFT of the x/nft module."),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"BurnNFT")," - Receive the user's request to burn an NFT, and call the BurnNFT of the x/nft module."),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"UpdateNFT")," - Receive the user's request to update an NFT, and call the UpdateNFT of the x/nft module.")),(0,l.kt)("h3",{id:"types-for-xnft-module"},"Types for x/nft Module"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"Class - describes the NFT class. Can be thought of as the smart contract address"),(0,l.kt)("li",{parentName:"ul"},"NFT - object representing unique non-fungible assets. Each NFT is associated with a Class")),(0,l.kt)("p",null,"Class Example:"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go="},"message Class {\n  string id          = 1;\n  string name        = 2;\n  string symbol      = 3;\n  string description = 4;\n  string uri         = 5;\n  string uri_hash    = 6;\n  google.protobuf.Any data = 7;\n}\n")),(0,l.kt)("p",null,"NFT Example:"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go="},"message NFT {\n  string class_id           = 1;\n  string id                 = 2;\n  string uri                = 3;\n  string uri_hash           = 4;\n  google.protobuf.Any data  = 10;\n}\n")),(0,l.kt)("h2",{id:"5-cw721-specification-in-detail"},"5. CW721 Specification In Detail"),(0,l.kt)("p",null,"The CW721 Specification will allow developers to create smart contracts based on CosmWasm. ",(0,l.kt)("em",{parentName:"p"},"Keep in mind all Messages will have a gas cost associated")),(0,l.kt)("h3",{id:"cw721-base-specfication"},"CW721-base Specfication"),(0,l.kt)("p",null,"InitMsg"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Takes the name and symbol for metadata, as well as the minter address. This is a special address that has full power to mint new NFTs but not modify existing NFTs.")),(0,l.kt)("p",null,"HandleMsg::Mint{token_id, owner, name, description, image}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Creates a new token with given owner and metadata. It can only be called by the Minter set in init.")),(0,l.kt)("p",null,"QueryMsg::Minter{}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Returns the minter address for this contract.")),(0,l.kt)("h2",{id:"base-messages"},"Base Messages"),(0,l.kt)("p",null,"TransferNFT{recipient, token_id}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Transfers ownership of the token to recipient account. Designed to send to an address and ",(0,l.kt)("em",{parentName:"p"},"does not")," trigger any actions to the recipient.")),(0,l.kt)("p",null,"SendNFT{contract, token_id, msg}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This transfers ownership of the token to the contract account specified in the message. The contract must be an address controlled by the contract, which implements the cw721Reciever interface. The msg will be passed along to the recipient contract along with the token_id.")),(0,l.kt)("p",null,"Approve{spender, token_id, expires}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Grants permission to spender to transfer or send the given token_id.")),(0,l.kt)("p",null,"Revoke{spender, token_id}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This revokes a previously granted permission to transfer the given token_id.")),(0,l.kt)("p",null,"ApproveAll{operator, expires}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Grants operator the permission to transfer or send all tokens owned by env.sender. This approval is tied to the owner, not the tokens and applies to any future token that the owner receives as well.")),(0,l.kt)("p",null,"RevokeAll{operator}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Revokes the previous ApproveAll permission granted to the given operator.")),(0,l.kt)("h3",{id:"base-queries"},"Base Queries"),(0,l.kt)("p",null,"OwnerOf{token_id}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Returns the owner of a given token, as well as anyone with approval on the particular token.")),(0,l.kt)("p",null,"ApprovedForAll{owner, include_expired}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"List all operators that can access all of the owner's tokens.")),(0,l.kt)("p",null,"NumToken{}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"Total number of tokens issued.")),(0,l.kt)("h3",{id:"base-receiver"},"Base Receiver"),(0,l.kt)("p",null,"ReceiveNft{sender, token_id, msg}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This is designed to handle SendNFT messages.")),(0,l.kt)("h3",{id:"metadata-queries"},"Metadata Queries"),(0,l.kt)("p",null,"ContractInfo{}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This returns top-level metadata about the contract. Namely the Name & symbol.")),(0,l.kt)("p",null,"NftInfo{token_id}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This returns metadata about one particular token. The return value is based on the ERC721 Metadata JSON Schema, but directly from the contract, not as a URI. Only the image link is a URI.")),(0,l.kt)("p",null,"AllNftInfo{token_id}"),(0,l.kt)("blockquote",null,(0,l.kt)("p",{parentName:"blockquote"},"This returns the result of both NftInfo and OwnerOf as one query as an optimization for clients, which may want both info to display one NFT.")))}h.isMDXComponent=!0}}]);