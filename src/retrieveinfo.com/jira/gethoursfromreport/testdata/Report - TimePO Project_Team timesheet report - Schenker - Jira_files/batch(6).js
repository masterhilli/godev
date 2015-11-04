;/* module-key = 'com.atlassian.plugins.atlassian-nav-links-plugin:rotp-menu', location = 'appswitcher/appswitcher.soy' */
// This file was automatically generated from appswitcher.soy.
// Please don't edit this file by hand.

if (typeof navlinks == 'undefined') { var navlinks = {}; }
if (typeof navlinks.templates == 'undefined') { navlinks.templates = {}; }
if (typeof navlinks.templates.appswitcher == 'undefined') { navlinks.templates.appswitcher = {}; }


navlinks.templates.appswitcher.linkSection = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  if (opt_data.list.length > 0) {
    output.append('<div class="aui-nav-heading sidebar-section-header">', soy.$$escapeHtml(opt_data.title), '</div><ul class="aui-nav nav-links">');
    var linkList8 = opt_data.list;
    var linkListLen8 = linkList8.length;
    for (var linkIndex8 = 0; linkIndex8 < linkListLen8; linkIndex8++) {
      var linkData8 = linkList8[linkIndex8];
      navlinks.templates.appswitcher.applicationsItem(linkData8, output);
    }
    output.append('</ul>');
  }
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.applicationsItem = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<li class="nav-link"><a href="', soy.$$escapeHtml(opt_data.link), '" class="interactive', (opt_data.self) ? ' checked' : '', '" title="', soy.$$escapeHtml(opt_data.link), '"><span class="nav-link-label">', soy.$$escapeHtml(opt_data.label), '</span></a></li>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.shortcutsItem = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<li class="nav-link"><a href="', soy.$$escapeHtml(opt_data.link), '" class="interactive', (opt_data.self) ? ' checked' : '', '" title="', soy.$$escapeHtml(opt_data.link), '"><span class="nav-link-label">', soy.$$escapeHtml(opt_data.label), '</span>', (opt_data.showDescription && opt_data.description) ? '<span class="nav-link-description">' + soy.$$escapeHtml(opt_data.description) + '</span>' : '', '</a></li>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.error = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<div class="app-switcher-error">', "Something went wrong, please \x3cspan class\x3d\x22app-switcher-retry\x22\x3etry again\x3c/span\x3e.", '</div>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.sidebarContents = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<div class="aui-page-panel-nav"><nav class="aui-navgroup aui-navgroup-vertical"><div class="app-switcher-section app-switcher-applications"><div class="aui-nav-heading">', soy.$$escapeHtml("Application Links"), '</div><div class="app-switcher-loading">', "Loading\x26hellip;", '</div></div><div class="app-switcher-section app-switcher-shortcuts"><div class="aui-nav-heading">', soy.$$escapeHtml("Shortcuts"), '</div><div class="app-switcher-loading">', "Loading\x26hellip;", '</div></div></nav></div>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.trigger = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<span class="aui-icon aui-icon-small aui-iconfont-appswitcher">', soy.$$escapeHtml("Linked Applications"), '</span>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.projectHeaderSection = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<div class="app-switcher-title">');
  aui.avatar.avatar({size: 'large', avatarImageUrl: opt_data.avatarUrl, isProject: true, title: opt_data.name}, output);
  output.append('<div class="sidebar-project-name">', soy.$$escapeHtml(opt_data.name), '</div></div>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.cogDropdown = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  var dropdownList__soy74 = new soy.StringBuilder();
  navlinks.templates.appswitcher.dropdownList({list: opt_data.links}, dropdownList__soy74);
  dropdownList__soy74 = dropdownList__soy74.toString();
  aui.dropdown2.dropdown2({menu: {'id': opt_data.id, 'content': dropdownList__soy74, 'extraClasses': 'aui-style-default sidebar-customize-section'}, trigger: {'showIcon': false, 'content': '<span class="aui-icon aui-icon-small aui-iconfont-configure"></span>', 'container': '#app-switcher'}}, output);
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.dropdownList = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<ul class="sidebar-admin-links">');
  var linkList82 = opt_data.list;
  var linkListLen82 = linkList82.length;
  for (var linkIndex82 = 0; linkIndex82 < linkListLen82; linkIndex82++) {
    var linkData82 = linkList82[linkIndex82];
    output.append('<li class="nav-link"><a href="', soy.$$escapeHtml(linkData82.href), '" title="', soy.$$escapeHtml(linkData82.title), '"><span class="nav-link-label">', soy.$$escapeHtml(linkData82.label), '</span></a></li>');
  }
  output.append('</ul>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.switcher = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  if (true) {
    if (AJS.DarkFeatures.isEnabled('rotp.sidebar')) {
      var sidebarContents__soy97 = new soy.StringBuilder();
      navlinks.templates.appswitcher.sidebarContents(null, sidebarContents__soy97);
      sidebarContents__soy97 = sidebarContents__soy97.toString();
      var triggerContent__soy99 = new soy.StringBuilder();
      navlinks.templates.appswitcher.trigger(null, triggerContent__soy99);
      triggerContent__soy99 = triggerContent__soy99.toString();
      navlinks.templates.appswitcher.sidebar({sidebar: {'id': 'app-switcher', 'content': sidebarContents__soy97}, trigger: {'showIcon': false, 'content': triggerContent__soy99}}, output);
      output.append('<script>\n                (function (NL) {\n                    var initialise = function () {\n                        new NL.SideBar({\n                            sidebarContents: \'#app-switcher\'\n                        });\n                    };\n                    if (NL.SideBar) {\n                        initialise();\n                    } else {\n                        NL.onInit = initialise;\n                    }\n                }(window.NL = (window.NL || {})));\n                window.NL.isUserAdmin = ', soy.$$escapeHtml(false), '<\/script>');
    } else {
      navlinks.templates.appswitcher_old.switcher(null, output);
    }
  }
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher.sidebar = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<a href="', soy.$$escapeHtml(opt_data.sidebar.id), '" class="sidebar-trigger app-switcher-trigger" aria-owns="', soy.$$escapeHtml(opt_data.sidebar.id), '" aria-haspopup="true">', opt_data.trigger.content, '</a><div id=', soy.$$escapeHtml(opt_data.sidebar.id), ' class="app-switcher-sidebar aui-style-default sidebar-offscreen">', opt_data.sidebar.content, '</div>');
  return opt_sb ? '' : output.toString();
};
;
;/* module-key = 'com.atlassian.plugins.atlassian-nav-links-plugin:rotp-menu', location = 'appswitcher/appswitcher.js' */
(function(c,a){a.SideBar=function(d){var e=this;this.$sidebar=null;d=c.extend({sidebarContents:null},d);this.getLinks=function(){return c.ajax({url:AJS.contextPath()+"/rest/menu/latest/appswitcher",cache:false,dataType:"json"}).done(this.updateAppLinks).fail(this.showAppSwitcherError)};this.populateProjectHeader=function(g,f){e.getSidebar().find(".app-switcher-shortcuts .aui-nav-heading").after(navlinks.templates.appswitcher.projectHeaderSection({avatarUrl:f,name:g}))};this.getProjectData=function(){var f=c(".project-shortcut-dialog-trigger"),g=f.data("key"),h=f.data("entity-type");if(f.size()==0||!g||!h){c(".app-switcher-shortcuts").remove();return}var j,i;i=c.ajax({url:AJS.contextPath()+"/rest/project-shortcuts/1.0/local/"+g,cache:false,data:{entityType:h},dataType:"json"});j=c.ajax({url:AJS.contextPath()+"/rest/project-shortcuts/1.0/remote/"+g,cache:false,data:{entityType:h},dataType:"json"});c.when(i,j).then(function(l,k){e.updateProjectShortcuts(l,k,{key:g,entityType:h,name:f.data("name"),avatarUrl:f.find("img").prop("src")})},e.showProjectShortcutsError)};this.getSidebar=function(){if(!this.$sidebar){this.$sidebar=c(d.sidebarContents)}return this.$sidebar};this.addApplicationsCog=function(){c(".app-switcher-applications .aui-nav-heading").before(navlinks.templates.appswitcher.cogDropdown({id:"sidebar-applications-admin-dropdown",links:[{href:AJS.contextPath()+"/plugins/servlet/customize-application-navigator",label:"Customize navigator",title:"Add new entries, hide existing or restrict who sees what"},{href:AJS.contextPath()+"/plugins/servlet/applinks/listApplicationLinks",label:"Manage application links",title:"Link to more Atlassian applications"}]}))};this.addProjectShortcutsCog=function(f,h){var g=[{href:AJS.contextPath()+"/plugins/servlet/custom-content-links-admin?entityKey="+f,label:"Customize shortcuts",title:""}];if(e.entityMappings[h]){g.push({href:e.generateEntityLinksUrl(f,e.entityMappings[h]),label:"Manage product links",title:""})}e.getSidebar().find(".app-switcher-shortcuts .aui-nav-heading").before(navlinks.templates.appswitcher.cogDropdown({id:"sidebar-project-shortcuts-admin-dropdown",links:g}))};this.updateAppLinks=function(f){c(function(){e.getSidebar().find(".app-switcher-applications").html(navlinks.templates.appswitcher.linkSection({title:"Application Links",list:f}));if(a.isUserAdmin){e.addApplicationsCog()}e.bindAnalyticsHandlers(e.getSidebar(),f)})};this.updateProjectShortcuts=function(i,g,h){var j=i[0].shortcuts,f=g[0].shortcuts;e.getSidebar().find(".app-switcher-shortcuts").html(navlinks.templates.appswitcher.linkSection({title:"Shortcuts",list:j.concat(f)}));if(a.isUserAdmin){e.addProjectShortcutsCog(h.key,h.entityType)}e.populateProjectHeader(h.name,h.avatarUrl);e.bindAnalyticsHandlers(e.getSidebar(),data)};this.entityMappings={"confluence.space":"com.atlassian.applinks.api.application.confluence.ConfluenceSpaceEntityType","jira.project":"com.atlassian.applinks.api.application.jira.JiraProjectEntityType","bamboo.project":"com.atlassian.applinks.api.application.bamboo.BambooProjectEntityType","stash.project":"com.atlassian.applinks.api.application.stash.StashProjectEntityType"};this.generateEntityLinksUrl=function(f,g){if(g===e.entityMappings["confluence.space"]){return AJS.contextPath()+"/spaces/listentitylinks.action?typeId="+g+"&key="+f}else{return AJS.contextPath()+"/plugins/servlet/applinks/listEntityLinks/"+g+"/"+f}};this.showAppSwitcherError=function(){c(function(){var f=e.getSidebar();f.find(".app-switcher-applications .app-switcher-loading").replaceWith(navlinks.templates.appswitcher.error());f.off(".appswitcher").on("click.appswitcher",".app-switcher-retry",c.proxy(e.retryLoading,e))})};this.showProjectShortcutsError=function(){c(function(){var f=e.getSidebar();f.find(".app-switcher-shortcuts .app-switcher-loading").replaceWith(navlinks.templates.appswitcher.error());f.off(".appswitcher").on("click.appswitcher",".app-switcher-retry",c.proxy(e.retryLoading,e))})};this.retryLoading=function(f){this.getSidebar().html(navlinks.templates.appswitcher.sidebarContents());this.getLinks();this.getProjectData();f&&f.stopPropagation()};this.bindAnalyticsHandlers=function(f,g){};this.getLinks();c(this.getProjectData);this.toggleSidebar=function(h){var i=e.getSidebar(),g=c("body"),f=c(window.document);if(!g.hasClass("app-switcher-open")){var k=c("#header");i.css("left",-i.width());i.parent("body").length||i.appendTo("body");b({data:i});i.animate({left:0},300);function j(l){var n=l.target&&c(l.target),m=l.keyCode;if(l.originalEvent===h.originalEvent){return}if(n&&!m&&!(n.closest(i).length||n.closest(k).length)&&h.which==1&&!(l.shiftKey||l.ctrlKey||l.metaKey)){e.toggleSidebar()}else{if(m===27){e.toggleSidebar()}}}f.on("click.appSwitcher",j);f.on("keydown.appSwitcher",j);f.on("scroll.appSwitcher",i,b)}else{f.off(".appSwitcher")}g.toggleClass("app-switcher-open")};c("#header").on("click",".app-switcher-trigger",this.toggleSidebar)};function b(f){var d=c(document).scrollTop(),g=c("#header"),e=(g.height()+g.offset().top)-d;if(e>=0){f.data.css({top:e,position:"fixed"})}else{f.data.css({top:0,left:0,position:"fixed"})}}if(a.onInit){a.onInit()}}(jQuery,window.NL=(window.NL||{})));;
;/* module-key = 'com.atlassian.plugins.atlassian-nav-links-plugin:rotp-menu', location = 'appswitcher/appswitcher_old.js' */
(function(e,c){c.AppSwitcher=function(f){var g=AJS.contextPath()+"/plugins/servlet/customize-application-navigator";var h=this;this.$dropdown=null;f=e.extend({dropdownContents:null},f);this.getLinks=function(){return e.ajax({url:AJS.contextPath()+"/rest/menu/latest/appswitcher",cache:false,dataType:"json"}).done(this.updateDropdown).fail(this.showError)};this.getDropdown=function(){if(!this.$dropdown){this.$dropdown=e(f.dropdownContents)}return this.$dropdown};this.updateDropdown=function(i){e(function(){h.getDropdown().html(navlinks.templates.appswitcher_old.applications({apps:i,showAdminLink:c.isUserAdmin,adminLink:g}));h.bindAnalyticsHandlers(h.getDropdown(),i);h.processSuggestionApps(i)})};this.bindAnalyticsHandlers=function(i,j){};this.processSuggestionApps=function(i){e.ajax({type:"GET",dataType:"json",contentType:"application/json; charset=UTF-8",url:AJS.contextPath()+"/rest/menu/latest/isAppSuggestionAvailable"}).done(function(j){if(j===true){h.handleSuggestionApps(i)}})};this.handleSuggestionApps=function(k){var l=_.map(k,function(m){return m.applicationType.toLowerCase()});$suggestionApps=e("<div id='app-switcher-suggestion-apps' class='aui-dropdown2-section'/>");$suggestionApps.html(navlinks.templates.appswitcher_old.suggestionApps);var j=$suggestionApps.find(".suggestion-apps");var i=false;_.each(a,function(m){if(!_.contains(l,m.appName)){i=true;j.append(navlinks.templates.appswitcher_old.suggestionApp({appName:m.appName,appDesc:m.appDesc}))}});if(!i){return}e("#app-switcher").append($suggestionApps);e(".app-discovery-suggestion-app").click(function(){AJS.trigger("analytics",{name:"appswitcher.discovery.user.select."+e(this).find("a").attr("id").toLowerCase()});window.open(e(this).find("a").attr("title"),"_blank")});e(".app-discovery-suggestion-app").hover(function(){e(this).find("a").removeClass("active").removeClass("aui-dropdown2-active")});e(".app-discovery-cancel-button").click(function(){AJS.trigger("analytics",{name:"appswitcher.discovery.nothanks.button.click"});d(b,"true");$suggestionApps.remove()})};this.showError=function(){e(function(){h.getDropdown().html(navlinks.templates.appswitcher_old.error()).off(".appswitcher").on("click.appswitcher",".app-switcher-retry",e.proxy(h.retryLoading,h))})};this.retryLoading=function(i){this.getDropdown().html(navlinks.templates.appswitcher_old.loading());this.getLinks();i&&i.stopPropagation()};this.getLinks()};var b="key-no-thanks";var a=[{appName:"jira",appDesc:"Issue & Project Tracking Software"},{appName:"confluence",appDesc:"Collaboration and content sharing"},{appName:"bamboo",appDesc:"Continuous integration"}];var d=function(f,g){e.ajax({url:AJS.contextPath()+"/rest/menu/latest/userdata/",type:"PUT",contentType:"application/json",data:JSON.stringify({key:f,value:g})})};if(c.onInit){c.onInit()}}(jQuery,window.NL=(window.NL||{})));;
;/* module-key = 'com.atlassian.plugins.atlassian-nav-links-plugin:rotp-menu', location = 'appswitcher/appswitcher_old.soy' */
// This file was automatically generated from appswitcher_old.soy.
// Please don't edit this file by hand.

if (typeof navlinks == 'undefined') { var navlinks = {}; }
if (typeof navlinks.templates == 'undefined') { navlinks.templates = {}; }
if (typeof navlinks.templates.appswitcher_old == 'undefined') { navlinks.templates.appswitcher_old = {}; }


navlinks.templates.appswitcher_old.applications = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  navlinks.templates.appswitcher_old.applicationsSection({list: opt_data.apps, listClass: 'nav-links', showDescription: opt_data.showDescription}, output);
  if (opt_data.custom) {
    navlinks.templates.appswitcher_old.applicationsSection({list: opt_data.custom, showDescription: opt_data.showDescription}, output);
  }
  if (opt_data.showAdminLink) {
    navlinks.templates.appswitcher_old.adminSection(opt_data, output);
  }
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.applicationsSection = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  if (opt_data.list.length > 0) {
    var param19 = new soy.StringBuilder('<ul', (opt_data.listClass) ? ' class="' + soy.$$escapeHtml(opt_data.listClass) + '"' : '', '>');
    var linkList27 = opt_data.list;
    var linkListLen27 = linkList27.length;
    for (var linkIndex27 = 0; linkIndex27 < linkListLen27; linkIndex27++) {
      var linkData27 = linkList27[linkIndex27];
      navlinks.templates.appswitcher_old.applicationsItem(soy.$$augmentData(linkData27, {showDescription: opt_data.showDescription}), param19);
    }
    param19.append('</ul>');
    aui.dropdown2.section({content: param19.toString()}, output);
  }
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.applicationsItem = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<li class="nav-link', (opt_data.self) ? ' nav-link-local' : '', '"><a href="', soy.$$escapeHtml(opt_data.link), '" class="aui-dropdown2-radio interactive', (opt_data.self) ? ' checked' : '', '" title="', soy.$$escapeHtml(opt_data.link), '"><span class="nav-link-label">', soy.$$escapeHtml(opt_data.label), '</span>', (opt_data.showDescription && opt_data.description) ? '<span class="nav-link-description">' + soy.$$escapeHtml(opt_data.description) + '</span>' : '', '</a></li>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.adminSection = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  aui.dropdown2.section({content: '<ul class="nav-links"><li><a class="nav-link-edit-wrapper" href="' + soy.$$escapeHtml(opt_data.adminLink) + '"><span class="nav-link-edit">' + "Configure\x26hellip;" + '</span></a></li></ul>'}, output);
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.error = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<div class="app-switcher-error">', "Something went wrong, please \x3cspan class\x3d\x22app-switcher-retry\x22\x3etry again\x3c/span\x3e.", '</div>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.loading = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<div class="app-switcher-loading">', "Loading\x26hellip;", '</div>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.trigger = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<span class="aui-icon aui-icon-small aui-iconfont-appswitcher">', soy.$$escapeHtml("Linked Applications"), '</span>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.switcher = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  if (true) {
    var loadingContent__soy81 = new soy.StringBuilder();
    navlinks.templates.appswitcher_old.loading(null, loadingContent__soy81);
    loadingContent__soy81 = loadingContent__soy81.toString();
    var triggerContent__soy83 = new soy.StringBuilder();
    navlinks.templates.appswitcher_old.trigger(null, triggerContent__soy83);
    triggerContent__soy83 = triggerContent__soy83.toString();
    aui.dropdown2.dropdown2({menu: {'id': 'app-switcher', 'content': loadingContent__soy81, 'extraClasses': 'aui-style-default'}, trigger: {'showIcon': false, 'content': triggerContent__soy83, 'extraClasses': 'app-switcher-trigger'}}, output);
    output.append('<script>\n            (function (NL) {\n                var initialise = function () {\n                    // For some milestones of AUI, the atlassian soy namespace was renamed to aui. Handle that here by ensuring that window.atlassian is defined.\n                    window.atlassian = window.atlassian || window.aui;\n                    new NL.AppSwitcher({\n                        dropdownContents: \'#app-switcher\'\n                    });\n                };\n                if (NL.AppSwitcher) {\n                    initialise();\n                } else {\n                    NL.onInit = initialise;\n                }\n            }(window.NL = (window.NL || {})));\n            window.NL.isUserAdmin = ', soy.$$escapeHtml(false), '<\/script>');
  }
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.suggestionApp = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<li class="app-discovery-suggestion-app"><a id="', soy.$$escapeHtml(opt_data.appName), '" href="#" class="app-discovery-link aui-icon-container app-discovery-', soy.$$escapeHtml(opt_data.appName), '-product-icon" title="https://www.atlassian.com/software/', soy.$$escapeHtml(opt_data.appName), '"/><div class="app-discovery-small">', soy.$$escapeHtml(opt_data.appDesc), '</div></li>');
  return opt_sb ? '' : output.toString();
};


navlinks.templates.appswitcher_old.suggestionApps = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<ul class=\'nav-links suggestion-apps\'><li><span class=\'app-discovery-suggest-title nav-link-label\'><h6>', soy.$$escapeHtml("Try other Atlassian apps"), '</h6></span></li></ul><div class=\'buttons-container app-discovery-suggest-apps-buttons\'><div class=\'buttons\'><button class=\'aui-button aui-button-link app-discovery-cancel-button\' name=\'cancel\' accesskey=\'c\' href=\'#\'>', soy.$$escapeHtml("Don\x27t show this again"), '</button></div></div>');
  return opt_sb ? '' : output.toString();
};
;
;/* module-key = 'com.atlassian.plugins.atlassian-nav-links-plugin:rotp-menu', location = 'appswitcher/inline-dialog.js' */
(function(d){AJS.CustomInlineDialog=AJS.CustomInlineDialog||function(B,n,r,p){p=p||[];if(p.hasOwnProperty("getArrowAttributes")){getArrowAttributesDeprecationLogger()}if(p.hasOwnProperty("getArrowPath")){getArrowPathDeprecationLogger();if(p.hasOwnProperty("gravity")){getArrowPathWithGravityDeprecationLogger()}}if(p.hasOwnProperty("onTop")){onTopDeprecationLogger();if(p.onTop&&p.gravity===undefined){p.gravity="s"}}if(typeof n==="undefined"){n=String(Math.random()).replace(".","");if(d("#inline-dialog-"+n+", #arrow-"+n+", #inline-dialog-shim-"+n).length){throw"GENERATED_IDENTIFIER_NOT_UNIQUE"}}var z=d.extend(false,AJS.CustomInlineDialog.opts,p);if(z.gravity==="w"){z.offsetX=p.offsetX===undefined?10:p.offsetX;z.offsetY=p.offsetY===undefined?0:p.offsetY}var v=function(){return window.Raphael&&p&&(p.getArrowPath||p.getArrowAttributes)};var j;var q;var K;var u=false;var A=false;var I=false;var J;var x;var g=d('<div id="inline-dialog-'+n+'" class="aui-inline-dialog"><div class="aui-inline-dialog-contents contents"></div><div id="arrow-'+n+'" class="aui-inline-dialog-arrow arrow"></div></div>');var m=d("#arrow-"+n,g);var G=g.find(".contents");if(!v()){g.find(".aui-inline-dialog-arrow").addClass("aui-css-arrow")}if(!z.displayShadow){G.addClass("aui-inline-dialog-no-shadow")}if(z.autoWidth){G.addClass("aui-inline-dialog-auto-width")}else{G.css("width",z.width+"px")}G.mouseover(function(O){clearTimeout(q);g.unbind("mouseover")}).mouseout(function(){E()});var D=function(){if(!j){j={popup:g,hide:function(){E(0)},id:n,show:function(){y()},persistent:z.persistent?true:false,reset:function(){function S(U,T){U.css(T.popupCss);if(v()){if(T.gravity==="s"){T.arrowCss.top-=d.browser.msie?10:9}if(!U.arrowCanvas){U.arrowCanvas=Raphael("arrow-"+n,16,16)}var V=z.getArrowPath,W=d.isFunction(V)?V(T):V;U.arrowCanvas.path(W).attr(z.getArrowAttributes())}else{m.removeClass("aui-bottom-arrow aui-left-arrow aui-right-arrow");if(T.gravity==="s"&&!m.hasClass("aui-bottom-arrow")){m.addClass("aui-bottom-arrow")}else{if(T.gravity==="n"){}else{if(T.gravity==="w"){m.addClass("aui-left-arrow")}else{if(T.gravity==="e"){m.addClass("aui-right-arrow")}}}}}m.css(T.arrowCss)}var P=AJS.$(window).height();var Q=Math.round(P*0.75);g.children(".aui-inline-dialog-contents").css("max-height",Q);var O=z.calculatePositions(g,x,J,z);if(O.hasOwnProperty("displayAbove")){displayAboveDeprecationLogger();O.gravity=O.displayAbove?"s":"n"}S(g,O);g.fadeIn(z.fadeTime,function(){});if(d.browser.msie&&~~(d.browser.version)<10){var R=d("#inline-dialog-shim-"+n);if(!R.length){d(g).prepend(d('<iframe class = "inline-dialog-shim" id="inline-dialog-shim-'+n+'" frameBorder="0" src="javascript:false;"></iframe>'))}R.css({width:G.outerWidth(),height:G.outerHeight()})}}}}return j};var y=function(){if(g.is(":visible")){return}K=setTimeout(function(){if(!I||!A){return}z.addActiveClass&&d(B).addClass("active");u=true;if(!z.persistent){H()}AJS.CustomInlineDialog.current=D();d(document).trigger("showLayer",["inlineDialog",D()]);D().reset()},z.showDelay)};var E=function(O){if(typeof O=="undefined"&&z.persistent){return}A=false;if(u&&z.preHideCallback.call(g[0].popup)){O=(O==null)?z.hideDelay:O;clearTimeout(q);clearTimeout(K);if(O!=null){q=setTimeout(function(){l();z.addActiveClass&&d(B).removeClass("active");g.fadeOut(z.fadeTime,function(){z.hideCallback.call(g[0].popup)});if(g.arrowCanvas){g.arrowCanvas.remove();g.arrowCanvas=null}u=false;A=false;d(document).trigger("hideLayer",["inlineDialog",D()]);AJS.CustomInlineDialog.current=null;if(!z.cacheContent){I=false;w=false}},O)}}};var F=function(R,P){var O=d(P);z.upfrontCallback.call({popup:g,hide:function(){E(0)},id:n,show:function(){y()}});g.each(function(){if(typeof this.popup!="undefined"){this.popup.hide()}});if(z.closeOthers){d(".aui-inline-dialog").each(function(){!this.popup.persistent&&this.popup.hide()})}x={target:O};if(!R){J={x:O.offset().left,y:O.offset().top}}else{J={x:R.pageX,y:R.pageY}}if(!u){clearTimeout(K)}A=true;var Q=function(){w=false;I=true;z.initCallback.call({popup:g,hide:function(){E(0)},id:n,show:function(){y()}});y()};if(!w){w=true;if(d.isFunction(r)){r(G,P,Q)}else{d.get(r,function(T,S,U){G.html(z.responseHandler(T,S,U));I=true;z.initCallback.call({popup:g,hide:function(){E(0)},id:n,show:function(){y()}});y()})}}clearTimeout(q);if(!u){y()}return false};g[0].popup=D();var w=false;var t=false;var s=function(){if(!t){d(z.container).append(g);t=true}};var o=d(B);if(z.onHover){if(z.useLiveEvents){if(o.selector){d(document).on("mousemove",o.selector,function(O){s();F(O,this)}).on("mouseout",o.selector,function(){E()})}else{AJS.log("Warning: inline dialog trigger elements must have a jQuery selector when the useLiveEvents option is enabled.")}}else{o.mousemove(function(O){s();F(O,this)}).mouseout(function(){E()})}}else{if(!z.noBind){if(z.useLiveEvents){if(o.selector){d(document).on("click",o.selector,function(O){s();if(M()){g.hide()}else{F(O,this)}return false}).on("mouseout",o.selector,function(){E()})}else{AJS.log("Warning: inline dialog trigger elements must have a jQuery selector when the useLiveEvents option is enabled.")}}else{o.click(function(O){s();if(M()){g.hide()}else{F(O,this)}return false}).mouseout(function(){E()})}}}var M=function(){return u&&z.closeOnTriggerClick};var H=function(){k();f()};var l=function(){C();N()};var i=false;var h=n+".inline-dialog-check";var k=function(){if(!i){d("body").bind("click."+h,function(P){var O=d(P.target);if(O.closest("#inline-dialog-"+n+" .contents").length===0){E(0)}});i=true}};var C=function(){if(i){d("body").unbind("click."+h)}i=false};var L=function(O){if(O.keyCode===27){E(0)}};var f=function(){d(document).on("keydown",L)};var N=function(){d(document).off("keydown",L)};g.show=function(P,O){if(P){P.stopPropagation()}s();if(z.noBind&&!(B&&B.length)){F(P,O===undefined?P.target:O)}else{F(P,B)}};g.hide=function(){E(0)};g.refresh=function(){if(u){D().reset()}};g.getOptions=function(){return z};return g};function c(g){var f=d(g);var h=d.extend({left:0,top:0},f.offset());return{left:h.left,top:h.top,width:f.outerWidth(),height:f.outerHeight()}}function b(h,j,t,f){var m=AJS.$.isFunction(f.offsetX)?f.offsetX(h,j,t,f):f.offsetX;var l=AJS.$.isFunction(f.offsetY)?f.offsetY(h,j,t,f):f.offsetY;var q=AJS.$.isFunction(f.arrowOffsetX)?f.arrowOffsetX(h,j,t,f):f.arrowOffsetX;var p=AJS.$.isFunction(f.arrowOffsetY)?f.arrowOffsetY(h,j,t,f):f.arrowOffsetY;var s=f.container.toLowerCase()!=="body";var g=AJS.$(f.container);var o=s?AJS.$(f.container).parent():AJS.$(window);var r=s?g.offset():{left:0,top:0};var k=s?o.offset():{left:0,top:0};var i=j.target;var u=i.offset();var n=i[0].getBBox&&i[0].getBBox();return{screenPadding:10,arrowMargin:5,window:{top:k.top,left:k.left,scrollTop:o.scrollTop(),scrollLeft:o.scrollLeft(),width:o.width(),height:o.height()},scrollContainer:{width:g.width(),height:g.height()},trigger:{top:u.top-r.top,left:u.left-r.left,width:n?n.width:i.outerWidth(),height:n?n.height:i.outerHeight()},dialog:{width:h.width(),height:h.height(),offset:{top:l,left:m}},arrow:{height:h.find(".arrow").outerHeight(),offset:{top:p,left:q}}}}function e(f,p,G,r){var n=b(f,p,G,r);var j=n.screenPadding;var k=n.window;var t=n.trigger;var D=n.dialog;var i=n.arrow;var z=n.scrollContainer;var E={top:t.top-k.scrollTop,left:t.left-k.scrollLeft};var A=Math.floor(t.height/2);var v=Math.floor(D.height/2);var u=Math.floor(i.height/2);var C=E.left-D.offset.left-j;var H=z.width-E.left-t.width-D.offset.left-j;var B=C>=D.width;var h=H>=D.width;var l=!h&&B?"e":"w";var o=E.top+A-u;var q=k.height-o-i.height;j=Math.min(j,o-n.arrowMargin);j=Math.min(j,q-n.arrowMargin);var g=E.top+A;var x=Math.max(g-j,0);var F=Math.max(k.height-g-j,0);var y=v-D.offset.top>x;var m=v+D.offset.top>F;var w;var s;if(y){w={top:k.scrollTop+j,left:l==="w"?t.left+t.width+D.offset.left:t.left-D.width-D.offset.left};s={top:(t.top+A)-(w.top+u)}}else{if(m){w={top:k.scrollTop+k.height-D.height-j,left:l==="w"?t.left+t.width+D.offset.left:t.left-D.width-D.offset.left};s={top:(t.top+A)-(w.top+u)}}else{w={top:t.top+A-v+D.offset.top,left:l==="w"?t.left+t.width+D.offset.left:t.left-D.width-D.offset.left};s={top:v-u+i.offset.top}}}return{gravity:l,popupCss:w,arrowCss:s}}function a(g,o,z,s){var x=AJS.$.isFunction(s.offsetX)?s.offsetX(g,o,z,s):s.offsetX;var v=AJS.$.isFunction(s.offsetY)?s.offsetY(g,o,z,s):s.offsetY;var m=AJS.$.isFunction(s.arrowOffsetX)?s.arrowOffsetX(g,o,z,s):s.arrowOffsetX;var l=AJS.$.isFunction(s.arrowOffsetY)?s.arrowOffsetY(g,o,z,s):s.arrowOffsetY;var f=c(window);var p=c(o.target);var y=c(g);var j=c(g.find(".aui-inline-dialog-arrow"));var i=p.left+p.width/2;var u=(window.pageYOffset||document.documentElement.scrollTop)+f.height;var k=10;y.top=p.top+p.height+~~v;y.left=p.left+~~x;var r=f.width-(y.left+y.width+k);j.left=i-y.left+~~m;j.top=-(j.height/2);var n=p.top>y.height;var h=(y.top+y.height)<u;var q=(!h&&n)||(n&&s.gravity==="s");if(q){y.top=p.top-y.height-(j.height/2);j.top=y.height}if(s.isRelativeToMouse){if(r<0){y.right=k;y.left="auto";j.left=z.x-(f.width-y.width)}else{y.left=z.x-20;j.left=z.x-y.left}}else{if(r<0){y.right=k;y.left="auto";var w=f.width-y.right;var t=w-y.width;j.right="auto";j.left=i-t-j.width/2}else{if(y.width<=p.width/2){j.left=y.width/2;y.left=i-y.width/2}}}return{gravity:q?"s":"n",displayAbove:q,popupCss:{left:y.left,top:y.top,right:y.right},arrowCss:{left:j.left,top:j.top,right:j.right}}}AJS.CustomInlineDialog.opts={onTop:false,responseHandler:function(g,f,h){return g},closeOthers:true,isRelativeToMouse:false,addActiveClass:true,onHover:false,useLiveEvents:false,noBind:false,fadeTime:100,persistent:false,hideDelay:10000,showDelay:0,width:300,offsetX:0,offsetY:10,arrowOffsetX:0,arrowOffsetY:0,container:"body",cacheContent:true,displayShadow:true,autoWidth:false,gravity:"n",closeOnTriggerClick:false,preHideCallback:function(){return true},hideCallback:function(){},initCallback:function(){},upfrontCallback:function(){},calculatePositions:function(f,i,j,h){h=h||{};var g=h.gravity==="w"?e:a;return g(f,i,j,h)},getArrowPath:function(f){if(f.gravity==="s"){return"M0,8L8,16,16,8"}else{return"M0,8L8,0,16,8"}},getArrowAttributes:function(){return{fill:"#fff",stroke:"#ccc"}}}})(AJS.$);;
;/* module-key = 'com.atlassian.jira.jira-header-plugin:jira-header', location = 'soy/headerDropdown.soy' */
// This file was automatically generated from headerDropdown.soy.
// Please don't edit this file by hand.

if (typeof JIRA == 'undefined') { var JIRA = {}; }
if (typeof JIRA.Templates == 'undefined') { JIRA.Templates = {}; }
if (typeof JIRA.Templates.Menu == 'undefined') { JIRA.Templates.Menu = {}; }
if (typeof JIRA.Templates.Menu.Dropdowns == 'undefined') { JIRA.Templates.Menu.Dropdowns = {}; }


JIRA.Templates.Menu.Dropdowns.dropdown2Fragment = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  var sectionList3 = opt_data.sections;
  var sectionListLen3 = sectionList3.length;
  for (var sectionIndex3 = 0; sectionIndex3 < sectionListLen3; sectionIndex3++) {
    var sectionData3 = sectionList3[sectionIndex3];
    var hasItems__soy4 = sectionData3.items && sectionData3.items.length > 0;
    output.append('<div class="aui-dropdown2-section">', (hasItems__soy4 && sectionData3.label) ? '<strong>' + soy.$$escapeHtml(sectionData3.label) + '</strong>' : '', '<ul class=\'aui-list-truncate\'', (sectionData3.id) ? ' id="' + soy.$$escapeHtml(sectionData3.id) + '"' : '', (sectionData3.style) ? ' class="' + soy.$$escapeHtml(sectionData3.style) + '"' : '', '>');
    if (hasItems__soy4) {
      var itemList25 = sectionData3.items;
      var itemListLen25 = itemList25.length;
      for (var itemIndex25 = 0; itemIndex25 < itemListLen25; itemIndex25++) {
        var itemData25 = itemList25[itemIndex25];
        JIRA.Templates.Menu.Dropdowns.dropdown2Item(itemData25, output);
      }
    }
    output.append('</ul></div>');
  }
  return opt_sb ? '' : output.toString();
};


JIRA.Templates.Menu.Dropdowns.dropdown2Item = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<li', (opt_data.id) ? ' id="' + soy.$$escapeHtml(opt_data.id) + '"' : '', (opt_data.style) ? ' class="' + soy.$$escapeHtml(opt_data.style) + '"' : '', '><a href="', soy.$$escapeHtml(opt_data.url), '"', (opt_data.id) ? ' id="' + soy.$$escapeHtml(opt_data.id) + '_lnk"' : '', (opt_data.title) ? ' title="' + soy.$$escapeHtml(opt_data.title) + '"' : '', ' class="', (opt_data.iconUrl) ? 'aui-icon-container' : '', (opt_data.parameters && opt_data.parameters['class']) ? ' ' + soy.$$escapeHtml(opt_data.parameters['class']) : '', '"');
  if (opt_data.parameters) {
    var keyList66 = soy.$$getMapKeys(opt_data.parameters);
    var keyListLen66 = keyList66.length;
    for (var keyIndex66 = 0; keyIndex66 < keyListLen66; keyIndex66++) {
      var keyData66 = keyList66[keyIndex66];
      output.append((keyData66 != 'class') ? ' ' + soy.$$escapeHtml(keyData66) + '="' + soy.$$escapeHtml(opt_data.parameters[keyData66]) + '"' : '');
    }
  }
  output.append('>', (opt_data.iconUrl) ? '<img class="icon" src="' + soy.$$escapeHtml(opt_data.iconUrl) + '" />' : '', soy.$$escapeHtml(opt_data.label), '</a></li>');
  return opt_sb ? '' : output.toString();
};
;
;/* module-key = 'com.atlassian.jira.jira-header-plugin:jira-header', location = 'js/init-dropdown2.js' */
;(function() {
    var SmartAjax = require('jira/ajs/ajax/smart-ajax');
    var Events = require('jira/util/events');
    var $ = require('jquery');

    /**
     * Binds analytics events to capture opening / closing of any dropdown menus, and clicking of any menu items therein.
     * Passes the id of the link element as a parameter
     */
    function bindHeaderDropdown2Analytics() {
        $("nav.aui-header a.aui-dropdown2-trigger").each(function() {
            var $ddtrigger = $(this);
            var id = $ddtrigger.attr("id");
            if (id) {
                var $dd = $("#" + $ddtrigger.attr("aria-owns"));
                $dd.on("aui-dropdown2-show", function() {
                    AJS.trigger("analytics", {
                        name: "jira.header.menu.opened",
                        data: {
                            id: id
                        }
                    });
                });
                $dd.on("aui-dropdown2-hide", function() {
                    AJS.trigger("analytics", {
                        name: "jira.header.menu.closed",
                        data: {
                            id: id
                        }
                    });
                });
                $dd.on("click", "a", function() {
                    AJS.trigger("analytics", {
                        name: "jira.header.menu.item.clicked",
                        data: {
                            id: this.id,
                            menu_id: id
                        }
                    });
                });
            }
        });
    }

    /**
     * Binds all the header implementations of Dropdown2. Including global nav and user profile.
     */
    function bindHeaderDropdown2() {
        $(".jira-ajax-menu").each(function () {
            var $ddtrigger = $(this);
            var $dd = $("#" + $ddtrigger.attr("aria-owns"));
            var $ajaxkey = $dd.data("aui-dropdown2-ajax-key");

            if ($ajaxkey) {
                $dd.bind("aui-dropdown2-show", function (event, options) {
                    Events.trigger('aui-dropdown2-show-before', event.target.id);
                    $dd.empty();
                    $dd.addClass("aui-dropdown2-loading");
                    SmartAjax.makeRequest({
                        url: contextPath + "/rest/api/1.0/menus/" + $ajaxkey,
                        data: {
                            inAdminMode: AJS.Meta.getBoolean("in-admin-mode")
                        },
                        dataType: "json",
                        cache: false,
                        success: function (data) {
                            $dd.removeClass("aui-dropdown2-loading");
                            $dd.html(JIRA.Templates.Menu.Dropdowns.dropdown2Fragment(data));

                            if (options && options.selectFirst) {
                                $dd.find("a:not(.disabled)").filter(":first").addClass("active")
                            }
                            Events.trigger('aui-dropdown2-show-after', event.target.id);
                        }
                    });
                });
            }
        });
    }

    $(function () {
        bindHeaderDropdown2();
        bindHeaderDropdown2Analytics();
    });

    AJS.namespace('JIRA.Dropdowns.bindHeaderDropdown2', null, bindHeaderDropdown2);
})();
;
;/* module-key = 'com.atlassian.plugins.helptips.jira-help-tips:common', location = 'js/HelpTip.js' */
(function(f){function a(){return false}function d(){return true}var b="jira-help-tip aui-help";var k=0,g=new Date().getTime();function h(m){m=""+(m||"");return m.replace(/\./g,"-")}function j(n,m){if(AJS.EventQueue&&m&&m.attributes.id){var p={};var o=h(m.attributes.id);var q="";if(m.attributes.eventPrefix){q=m.attributes.eventPrefix;if(q.charAt(q.length-1)!=="."){q+="."}}p.name=q+"helptips."+o+"."+n;p.properties={};AJS.EventQueue.push(p)}}function c(){return"c"+g+(k++)}var e=AJS.HelpTip=function(m){var n;this.attributes=f.extend({},m);this.attributes.id||(this.attributes.id=false);this.attributes.callbacks||(this.attributes.callbacks={});if(this.attributes.isSequence){if(!this.attributes.weight){this.attributes.weight=Number.MAX_VALUE}AJS.HelpTip.Manager.sequences.push(this)}if(this.attributes.body){this.attributes.bodyHtml=this.attributes.body;delete this.attributes.body}this.cid=c();n=this.attributes.anchor;delete this.attributes.anchor;this.view=(n)?new i(this,n):new l(this)};f.extend(e.prototype,{show:function(o){o=o||{};var n=this;var p=f.Deferred();if(this.attributes.callbacks.beforeShow){var m=this.attributes.callbacks.beforeShow();if(m&&_.isFunction(m.done)){m.done(p.resolve)}else{p.resolve()}}else{p.resolve()}p.done(function(){AJS.HelpTip.Manager.show(function(){if(!n.isDismissed()){if(!o.force&&AJS.Popups&&AJS.Popups.DisplayController){AJS.Popups.DisplayController.request({name:n.id,weight:1000,show:function(){n.view.show()}})}else{n.view.show()}j("shown",n)}})})},dismiss:function(){var m=h(arguments[0]||"programmatically");this.view.dismiss();if(m==="close-button"&&this.attributes.isSequence){AJS.HelpTip.Manager.clearSequences()}if(!this.isDismissed()){AJS.HelpTip.Manager.dismiss(this);j("dismissed."+m,this)}},isVisible:function(){return this.view.$el.is(":visible")},isDismissed:function(){return AJS.HelpTip.Manager.isDismissed(this)},refresh:function(){if(!this.isDismissed()){this.view.refresh()}},hide:function(){if(!this.isDismissed()){this.view.dismiss()}},showNextHelpTipInSequence:function(){this.view.clickNext()}});var i=function(n,m){this.initialize(n,m)};f.extend(i.prototype,{initialize:function(n,m){this.model=n;this.anchorSelector=m;this.anchor=AJS.$(m);this._initDialog(m);AJS.$(document).bind("showLayer",function(q,p,o){if(p==="inlineDialog"&&o.id===n.cid){AJS.InlineDialog.current=null;AJS.$(document.body).unbind("click."+n.cid+".inline-dialog-check");o._validateClickToClose=a;o.hide=a}})},show:function(){this.beforeHide=a;this.popup.show()},refresh:function(){var m=AJS.$(this.anchorSelector);if(!m.is(":visible")){this.dismiss()}else{if(m.get(0)!==this.anchor.get(0)){this.changeAnchor(m)}else{if(!this.isVisible()){this.show()}else{this.popup.refresh()}}}},changeAnchor:function(m){var n=this.isVisible();this.dismiss();this.$el.remove();this.anchor=m;this._initDialog(m);if(n){this.show()}},dismiss:function(){this.beforeHide=d;this._popupHide()},clickNext:function(){var m=AJS.$(this.$el).find(".helptip-next");if(m.length>0){m.click()}},isVisible:function(){return this.$el.is(":visible")},_initDialog:function(o){var m=this;var n=this.model;this.popup=AJS.InlineDialog(f(o),n.cid,_.bind(this._createDialog,this),_.extend({container:"#content",noBind:true,preHideCallback:function(){return m.beforeHide()},calculatePositions:function(p,s,v,r){var q=AJS.InlineDialog.opts.calculatePositions(p,s,v,r);var u=f(this.container);var t=u.offset();if(q.popupCss.left!=="auto"){q.popupCss.left-=t.left;q.popupCss.right="auto"}q.popupCss.top-=t.top;return q},addActiveClass:false,initCallback:n.attributes.callbacks.init,hideCallback:n.attributes.callbacks.hide,persistent:true},n.attributes.inlineDialogOpts));this._popupHide=this.popup.hide;this.popup.hide=a;this.$el=f(this.popup[0]);this.$el.addClass(b)},_createDialog:function(r,p,o){var n=this;var q=AJS.HelpTip.Manager.sequences;var m=this.model.attributes.position;r.removeClass("contents");r.html(f(AJS.Templates.HelpTip.tipContent(_.extend({hasSequence:(q.length>1&&(m+1<q.length)),length:q.length,position:m,showCloseButton:true},this.model.attributes))));r.unbind("mouseover mouseout");r.find(".helptip-link").click(function(){j("learn-more.clicked",n.model)});r.find(".helptip-close").click(function(s){s.preventDefault();n.model.dismiss("close-button")});r.find(".helptip-next").click(function(t){t.preventDefault();n.model.dismiss("next-button");var s=m+1;q[s]&&(q[s].show({force:true}))});o()}});var l=function(m){this.initialize(m)};f.extend(l.prototype,{initialize:function(){this.$el=f("<div></div>");this.$el.addClass(b)},show:function(){},dismiss:function(){}})})(AJS.$);;
;/* module-key = 'com.atlassian.plugins.helptips.jira-help-tips:common', location = 'js/HelpTipManager.js' */
(function(d){var b=AJS.contextPath()+"/rest/helptips/1.0/tips";var c=undefined;if(WRM&&WRM.data){c=WRM.data.claim("com.atlassian.plugins.helptips.jira-help-tips:common.JiraHelpTipData")}var a={dismissedTipIds:[],sequences:[],loaded:d.Deferred(),url:function(){return b},sync:function(g,f){var e=d.Deferred();g||(g="get");f||(f=null);if(g==="get"&&c&&c.dismissed){e.resolve(c.dismissed)}else{d.ajax(this.url(),{type:g,dataType:"json",contentType:"application/json",data:f&&JSON.stringify(f),processData:false}).done(function(h){e.resolve(h)}).fail(function(){e.reject()})}return e.promise()},fetch:function(){var e=this.sync();e.done(d.proxy(function(f){d.merge(this.dismissedTipIds,f);this.loaded.resolve()},this));return e.promise()},show:function(e){this.loaded.done(e)},dismiss:function(e){var f=e.attributes.id;if(!f){e._dismissed=true}else{this.dismissedTipIds.push(f);this.sync("post",{id:f})}},undismiss:function(e){var f=e.attributes.id;if(!f){e._dismissed=false}else{this.dismissedTipIds.splice(d.inArray(f,this.dismissedTipIds),1);this.sync("delete",{id:f})}},isDismissed:function(e){var f=e.attributes.id;return(f)?d.inArray(f,this.dismissedTipIds)>=0:e._dismissed},clearSequences:function(){this.sequences=[]},hideSequences:function(){_.each(this.sequences,function(e){e.view.dismiss()})},showSequences:function(){if(!this._showStarted){var e=this;var f=0;this._showStarted=true;d.when(this.loaded).done(function(){e.sequences.sort(function(h,g){return h.attributes.weight-g.attributes.weight});e.sequences=_.filter(e.sequences,function(h){var g=_.indexOf(e.dismissedTipIds,h.attributes.id)===-1;if(g){h.attributes.position=f++}return g});if(e.sequences.length>0){e.sequences[0].show({force:true})}e._showStarted=false})}}};if(!JIRA.Users.LoggedInUser.isAnonymous()){AJS.HelpTip.Manager=a;a.fetch()}})(AJS.$);;
;/* module-key = 'com.atlassian.plugins.helptips.jira-help-tips:common', location = 'templates/HelpTip.soy' */
// This file was automatically generated from HelpTip.soy.
// Please don't edit this file by hand.

if (typeof AJS == 'undefined') { var AJS = {}; }
if (typeof AJS.Templates == 'undefined') { AJS.Templates = {}; }
if (typeof AJS.Templates.HelpTip == 'undefined') { AJS.Templates.HelpTip = {}; }


AJS.Templates.HelpTip.tipContent = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append((opt_data.title) ? '<h2 class="helptip-title">' + soy.$$escapeHtml(opt_data.title) + '</h2>' : '', '<p class="helptip-body">', opt_data.bodyHtml, '</p>', (opt_data.url) ? '<p><a class="helptip-link" href="' + soy.$$escapeHtml(opt_data.url) + '" target="_blank">' + ((opt_data.linkText) ? soy.$$escapeHtml(opt_data.linkText) : soy.$$escapeHtml("Learn more")) + '</a></p>' : '');
  AJS.Templates.HelpTip.tipFooter(opt_data, output);
  return opt_sb ? '' : output.toString();
};


AJS.Templates.HelpTip.tipFooter = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append('<form class="tip-footer">');
  AJS.Templates.HelpTip.nextButton(opt_data, output);
  AJS.Templates.HelpTip.closeButton(opt_data, output);
  AJS.Templates.HelpTip.sequencePaging(opt_data, output);
  output.append('</form>');
  return opt_sb ? '' : output.toString();
};


AJS.Templates.HelpTip.nextButton = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append((opt_data.hasSequence) ? '<button class="aui-button helptip-next" type="button">' + ((opt_data.nextButtonText) ? soy.$$escapeHtml(opt_data.nextButtonText) : soy.$$escapeHtml("Next")) + '</button>' : '');
  return opt_sb ? '' : output.toString();
};


AJS.Templates.HelpTip.closeButton = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append((opt_data.showCloseButton) ? '<button class="aui-button ' + ((opt_data.hasSequence) ? ' aui-button-link ' : '') + ' helptip-close" type="button">' + ((opt_data.closeButtonText) ? soy.$$escapeHtml(opt_data.closeButtonText) : soy.$$escapeHtml("Close")) + '</button>' : '');
  return opt_sb ? '' : output.toString();
};


AJS.Templates.HelpTip.sequencePaging = function(opt_data, opt_sb) {
  var output = opt_sb || new soy.StringBuilder();
  output.append((opt_data.length > 1) ? '<span class="helptip-sequence-paging">' + soy.$$escapeHtml(opt_data.position + 1) + '/' + soy.$$escapeHtml(opt_data.length) + '</span>' : '');
  return opt_sb ? '' : output.toString();
};
;
