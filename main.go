package main

import (
	"climatewhopper/pkg/newsparser"
	"fmt"

	"github.com/foolin/pagser"
)

func main() {
	p := pagser.New()
	website, _ := newsparser.GetWebsiteData("https://taz.de/Koalitionsverhandlungen-in-Berlin/!5816465/")
	resp, err := newsparser.TazParser.ParseArticle(p, &website)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.OriginalText)
}

var exampletext = `<html xmlns="http://www.w3.org/1999/xhtml" xmlns:my="mynames" lang="de"><!-- DEBUG start 21:53:31+01:00 page_id=4649 :: Berlin--><!--
Content Management: openNewspaper www.opennewspaper.org based on TYPO3 www.typo3.org
Community Platform: Invsision Power Board www.invisionpower.com via ipbwi.com + manufactured PHP
Presentation Layer: XML, XSL, HTML, CSS, JS (+ toil, tears and sweat) webmaster@taz.de
--><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<!--
	page_id :: 4649--><meta http-equiv="Content-Script-Type" content="text/javascript"><meta name="robots" content="index,follow,noarchive"><meta name="language" content="de"><meta name="copyright" content="TAZ Verlags- und Vertriebs GmbH"><title>Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung - taz.de</title><meta name="author" content="Claudius Prößer"><meta name="description" content="Die Neuauflage von Rot-Grün-Rot betont den Klimaschutz, will aber keine neuen Maßstäbe setzen: Ein Zeitplan, wann Berlin klimaneutral sein soll, fehlt."><meta name="keywords" content="Klimaschutz, R2G Berlin, Koalitionsverhandlungen, Kohleausstieg, Berlin, taz, tageszeitung "><meta name="taz:title" data-id="5816465" content="Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung"><meta name="taz:tag" content="Klimaschutz" data-tag-id="5009370"><meta name="taz:tag" content="R2G Berlin" data-tag-id="5364089"><meta name="taz:tag" content="Koalitionsverhandlungen" data-tag-id="5013056"><meta name="taz:tag" content="Kohleausstieg" data-tag-id="5204208"><meta name="generator" content="tazxslt, 0.55; "><meta property="fb:pages" content="171844246207985,162775943776229,136860635948,337939896245563,185657208283671,1141409319297464,255160261177600,1252816681448318,240508363106685,669240916596907"><meta property="og:url" content="https://taz.de/!5816465/"><meta property="og:locale" content="de_DE"><meta property="og:type" content="article"><meta name="twitter:card" content="summary_large_image"><meta name="twitter:site" content="@tazgezwitscher"><meta property="og:title" content="Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung"><meta property="fb:app_id" content="817458245779898"><meta property="article:opinion" content="false"><meta property="article:location" content="city:Berlin"><meta property="article:content_tier" content="free"><meta name="twitter:title" content="Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung"><meta property="og:description" content="Die Neuauflage von Rot-Grün-Rot betont den Klimaschutz, will aber keine neuen Maßstäbe setzen: Ein Zeitplan, wann Berlin klimaneutral sein soll, fehlt."><meta name="twitter:description" content=" Die Neuauflage von Rot-Grün-Rot betont den Klimaschutz, will aber keine neuen Maßstäbe setzen: Ein Zeitplan, wann Berlin klimaneutral sein soll, fehlt."><meta property="og:image" content="https://taz.de/picture/5230729/948/9770710-1.jpeg"><meta name="twitter:image" content="https://taz.de/picture/5230729/948/9770710-1.jpeg"><meta property="article:published_time" content="2021-11-18T11:10:00+01:00"><meta property="article:modified_time" content="2021-11-18T14:12:39+01:00">
<!-- piwik/matomo (0.6) --><script type="text/javascript" id="piwik-loaded" async="" defer="" src="https://taz.de/lib/share/js/piwik.js"></script><script id="piwik-init" type="text/javascript">

// ====================================================== taz piwik library === //
window.patOtaz_de = window.patOtaz_de || new Object;
window.patOtaz_de.piwik = window.patOtaz_de.piwik || new (function (debug) {

// for a better readability
Array.prototype.clone = Array.prototype.slice;

// ES6 features
var hasHistoryApi = function() {
if ( history && history.replaceState instanceof Function )  return true;
return false;
};

// === private static members
var version = "patOtaz_de.piwik-0.6",
debug   =  debug || false;

// === logging in debug mode or from outside
var log = function(args) {
if ( debug && arguments.length ) {
	var items = [];
	Array.prototype.push.apply( items, arguments );
	console.log( "PIWIK-LOG:", items );
}
return true;
};

// === clean up strings with a whitelist
var removeCharsWithWhitelist = function(validChars, inputString) {
var regex = new RegExp( '[^'+ validChars +']', 'g' );
return inputString.replace( regex, '' );
};

// === clean up #matomo url fragment and return matomo part
var removeHashParam = function(fragment_key) {
if ( fragment_key === undefined )  return '';
var href          = location.href;
var href_base     = href.split('#')[0];
var fragment_pos  = href.indexOf('#');
var fragment_full = href.substring( fragment_pos +1 );
if ( fragment_pos >0  &&  fragment_key.length >0 ) {
	var fragment_regex = new RegExp( '(.*?)#?'+ fragment_key +':([^#]+)#?(.*)' );
	var fragment_match = fragment_full.match( fragment_regex );
	if ( fragment_match ) {  
		var fragment_this    = fragment_match[2];
		var fragment_lhs     = fragment_match[1];
		var fragment_rhs     = fragment_match[3];
		var fragment_cleaned = ( fragment_lhs ? fragment_lhs : '' )+( fragment_lhs && fragment_rhs ? '#' : '' )+( fragment_rhs ? fragment_rhs : '' );
		if ( fragment_this.length >0 ) {
			var new_href = href_base +( !fragment_lhs && !fragment_rhs ? '' : '#' )+ fragment_cleaned;
			if ( hasHistoryApi )  history.replaceState( {}, document.title, new_href );
			return fragment_this;
		}
	}
}
return '';
};

// === track search results list page
var trackSiteSearch = function(keyword, category, count) {
var tmp = Array.prototype.clone.call( arguments );
tmp.unshift("trackSiteSearch");
_paq.push( tmp );
log( "search result", { keyword:keyword, category:category, count:count } );
return true;
};

// === track funnels
var trackFunnel = function(funnel, category) {
if ( funnel.length ) {
	_paq.push([ 'trackEvent', category, 'Funnel', funnel ]);
	log( "goal event", { funnel:funnel, category:category } );
}
return true;
};

// === track events
var trackEvent = function(category, action, name, value) {
var tmp = Array.prototype.clone.call( arguments );
tmp.unshift("trackEvent");
_paq.push( tmp );
log( "event", { category:category, action:action, name:name, value:value } );
return true;
}

// === debugging
log( version );

// === public interface
this.log              =log;
this.removeHashParam  =removeHashParam;
this.trackSiteSearch  =trackSiteSearch;
this.trackFunnel      =trackFunnel;
this.trackEvent       =trackEvent;
// this.trackPageType    =trackPageType;

})(false); // patOtaz_de.piwik


// ====================================================== piwik bucket === //
window._paq = window._paq || [];


// ====================================================== piwik setup === //
(function(){

// === private members
var p = patOtaz_de.piwik,
o = {
	version               :"0.6" // String
	, piwik_script        :"https://taz.de/stats/piwik.php" // String
	, page_id_enc         :"5816465" // String
	, page_title_enc      :"5816465: Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung" // String
	, custom_url_enc      :"/!5816465" // String
	, channel             :"web" // String
	, area                :"Redaktion" // String
	, department          :"p4649" // String
	, piwik_site_id       : 1  // Number
	, is_tag              : false  // Boolean
	, is_article          : true  // Boolean
	, has_tags            : true  // Boolean
};

// === url decode strings
o.page_id         =decodeURIComponent( o.page_id_enc    );
o.page_title      =decodeURIComponent( o.page_title_enc );
o.custom_url      =decodeURIComponent( o.custom_url_enc );

// === page title
!function( node ) {
try {
	var id    = node.getAttribute('data-id').toString()
		, title = node.getAttribute('content').toString()
		;
	o.title = ( id.length >0 && title.length >0 ) ? id +": "+ title : "ZOMBIE";
} catch (err) {
	o.title = "ZOMBIE";
	console.log("ERROR-piwik: page title");
}
} ( document.head.querySelector( 'meta[ name="taz:title" ]' ) );

// === page url
o.url = location.href;

// === process and clean #matomo fragment, if present
o.url_fragment = decodeURIComponent( p.removeHashParam('matomo') );
if ( o.url_fragment.length >0 )  o.custom_url += '#'+ o.url_fragment;

// === track tags
if ( o.is_article && o.has_tags ) {
o.tag_list=[];
o.tag_list.toString = function() {
	var string="";
	// Get tag title and concat. So that no permutations are stored for the same set of tags, pull tag titles sorted by their id.
	this.sort( function(a,b){ return ( parseInt(a.id) - parseInt(b.id) ) } ).forEach( function(tag){ string += tag.title +";" });
	return string;
};
!function( tag_list ) {
	var count=0;
	try {
		tag_list.forEach( function(tag) {
			o.tag_list.push({
				title     :tag.getAttribute('content')
				,id       :tag.getAttribute('data-tag-id')
				,category :tag.getAttribute('data-tag-category')? tag.getAttribute('data-tag-category'): "PLAIN-VANILLA"
			});
			count++;
		});
		if (count===0)  o.tag_list.push({ title:'ZOMBIE', id:0, category:'ZOMBIE' });
		o.tag_list.string = o.tag_list.toString();
	} catch(err) {
		o.tag_list.push({ title:'ZOMBIE', id:0, category:'ZOMBIE' });
		console.log("ERROR-piwik: tag_list");
	}
}( document.head.querySelectorAll("meta[name='taz:tag']") );
}

// === setup piwik bucket
_paq.push([ "setCustomUrl"       ,o.custom_url ]);
_paq.push([ "setDocumentTitle"   ,o.page_title ]);
_paq.push([ "setCustomVariable"  ,1 ,"channel"    ,o.channel    ,"page" ]);
_paq.push([ "setCustomVariable"  ,2 ,"area"       ,o.area       ,"page" ]);
_paq.push([ "setCustomVariable"  ,3 ,"department" ,o.department ,"page" ]);
if ( o.is_article && o.has_tags )  _paq.push([ "setCustomVariable", 4, "tag", o.tag_list.toString(), "page" ]);
_paq.push([ "trackPageView"      ]);
_paq.push([ "enableLinkTracking" ]);

// === load and setup piwik tracking
(function() {
var d=document
  , g=d.createElement("script")
  , s=d.getElementById("piwik-init")
  ;
_paq.push([ "setTrackerUrl" ,o.piwik_script  ]);
_paq.push([ "setSiteId"     ,o.piwik_site_id ]);
g.type   = "text/javascript";
g.id     = "piwik-loaded";
g.async  = true;
g.defer  = true;
g.src    = "https://taz.de/lib/share/js/piwik.js";
g.onload = function() { p.log( 'setup', o ) }
s.parentNode.insertBefore( g, s );
})();

// === tag tracking
if ( o.is_article && o.has_tags ) {
o.tag_list.forEach( function(tag) {    p.trackEvent( "TAG", "TAG-ARTICLE-"+ tag.category, "TAG - t"+ tag.id +": "+ tag.title ) });
}
if ( o.is_tag ) {
p.trackEvent( "TAG", "TAG-LANDINGPAGE", "TAG - "+ o.title );
}

// === debugging
//p.log( _paq );

})(); // piwik setup

</script><!-- end: piwik/matomo (0.6) -->
<script type="text/javascript"> 
   
(function(){


// # fil 2021-09-16 # Intention? sieht aus, als würde es die Redirect-URL matchen, die kann ja aber nicht location.href sein?
if ( ((location.href.match(/goMobile/g) || []).length) > 1 ) {
console.log( 'prevent loop' );
return
}


if ( readCookie('ratioURL_channel') == 'web' ){ 
 return;
// web view was selected explicitly
}

// var getWidth = (screen.width > screen.height) ? screen.height : screen.width;	// # fil 2021-09-16 # klarer?
var getWidth;
if (screen.width > screen.height){
   getWidth = screen.height;
} else {
   getWidth = screen.width;
}


if (getWidth < 551) {
//mobile
var goto = location.pathname;
if (  (/;web/).test( location.href ) ){ 
	//prevent loop. stay desktop  
	return false;

}

// to do: look if this is running	# fil 2021-09-16 # versteh ich nicht.
// # fil 2021-09-16 # Intention? sieht aus, als würde es die Redirect-URL matchen, die kann ja aber nicht location.href etc sein?
if ( (/goMobile/).test( location.search ) && (/count/).test( location.href ) || ( location.search.indexOf('goMobile') > 1 ) ){	// back here despite cache-killer: give up.
		
		if ((/;web/).test( location.pathname )) {
			location.href = goto
			// # fil 2021-09-16 # setzt location.href=location.pathname - wozu??? Loop-Gefahr!

		};
		return false;
	}

if( (/moby/).test( readCookie('ratioURL_channel') ) ) { // back here despite cookie: reload w/ cache-killer to prevent loop
// Google-Bot needs this to see mobile (https://search.google.com/search-console/inspect fil 2018-12-03)
location.href = '/count/redirect/go-mobile-nocache'+ goto 
	+( location.search ? location.search +'&' : '?' ) +'goMobile2='+ (new Date()).getTime()
	+location.hash;

	// # fil 2021-09-16 # wird serverseitig nach goto+location.search umgeleitet, ?goMobile2=… wird dabei entfernt. Sinnvoll?
	return;
	}

// try to set channel cookie …
document.cookie = 'ratioURL_channel=moby; domain=taz.de; path=/';

if( (/moby/).test( readCookie('ratioURL_channel') ) ) {	// … cookie was successfully set: reload …
	if (location.search.indexOf('nocache') != -1) return;	// old cache killer for transition period
	location.reload( true );
	return;
	}
else {						// … setting cookie failed: load explicit channel url …
	// # fil 2021-09-16 # besser?: Cookie bei redirect serverseitig setzen?
	if (location.search.indexOf('nocache') != -1) return;	// old cache killer for transition period
	if( goto =='/' )
		goto = '/!p4608;moby/';
	else	goto = goto.replace(/[/]$/,';moby/');
	location.href = '/count/redirect/go-mobile-nocookie'+ goto 
		+location.search
		+location.hash;
	return;
	}

} else {
//desktop
}	

// vanilla js cookie reader
function readCookie(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for(var i=0;i < ca.length;i++) {
		var c = ca[i];
		while (c.charAt(0)==' ') c = c.substring(1,c.length);
		if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
	}
	return null;
}


})();	

</script><link rel="preload" href="//data-2d3a3249cd.taz.de/iomm/latest/manager/base/es6/bundle.js" as="script" id="IOMmBundle"><link rel="preload" href="//data-2d3a3249cd.taz.de/iomm/latest/bootstrap/loader.js" as="script"><script type="text/javascript" src="//data-2d3a3249cd.taz.de/iomm/latest/bootstrap/loader.js"></script><script src="https://data-2d3a3249cd.taz.de/iomm/latest/manager/base/es6/bundle.js"></script><link rel="preload" href="/lib/share/fonts/DroidSerif-Regular-webfont.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/AktivGroteskBold/AktivGrotesk_W_Bd.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/fontawesome-reduced.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_5_.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_6_.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_7_.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_8_.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_4_italic.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz_4_.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Taz-Bold_tazze_private_only.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/share/fonts/Quodana.woff2" type="font/woff2" crossorigin="anonymous" as="font"><link rel="preload" href="/lib/ch/moby/pix/tazze46.png" as="image" media="(min-width: 81px)"><link rel="preload" href="/lib/ch/web/pix/tzi_logo_120px.png" as="image" media="(min-width: 120px)"><link rel="preload" href="/lib/ch/web/pix/tazze_30_d50d2e.png" as="image" media="(min-width: 20px)"><link rel="canonical" href="https://taz.de/Koalitionsverhandlungen-in-Berlin/!5816465/"><link rel="alternate" type="application/rss+xml" title="taz.de - Berlin" href="/!p4649/rss.xml"><link rel="alternate" type="application/rss+xml" title="taz.de - Schlagzeilen" href="/rss.xml"><link rel="home" type="text/html" title="taz.de - Schlagzeilen" href="/"><link rel="copyright" type="text/html" title="Impressum" href="/6/impressum/"><link rel="stylesheet" type="text/css" media="screen, print" href="/lib/ch/web/css/news2020-12-08_20.css"><link href="/lib/ch/web/css/local/print.css" type="text/css" rel="stylesheet" media="print"><!--[if lte IE 8]><style type="text/css">
	p.article,
	.sect_article >.sectbody >h6,
	p.caption,
	.price-tag >.info,
	.ad_badge,
	li.tag >a,
	.secthead > ul.toolbar,
	.sect_meta,
	.sect_service >.sectbody,
	.head >.search >.frame >input,
	.person >h5 {
	  font-family: Verdana, DejaVu Sans, Bitstream Vera Sans, Helvetica, sans-serif;
  }
</style><![endif]--><script type="text/javascript" src="/lib/ch/moby/js/local/cmp.js"></script><script>

	if(! HTMLCollection.prototype.last ) HTMLCollection.prototype.last = function(){ return this[ this.length -1 ]; }; 
	if(! document.getLatestElement ) document.getLatestElement = function(){ 
		var scripts = document.getElementsByTagName("script"); var thisScript = scripts[scripts.length - 1];
		return ( this.currentScript || thisScript ).previousSibling;
		};

</script><script type="text/javascript" src="/lib/ch/web/js/2020-12-08_20.js"></script><!-- piwik/matomo (custom-0.2) --><script class="piwik-custom" type="text/javascript">

// ====================================================== piwik customizing === //
(function(){

// === private members
var p = patOtaz_de.piwik,
o = {
	version           :"custom-0.2" // String
	, page_id_enc     :"5816465" // String
	, form_funnel     :"" // String
	, search_term     :"" // String
	, rootline        :"/Berlin" // String
	, is_article      : true  // Boolean
	, is_corp         : false  // Boolean
	, is_searchresult : false  // Boolean
	, is_form         : false  // Boolean
	, search_count    : NaN  // Number
};

// === url decode strings
o.page_id = decodeURIComponent( o.page_id_enc );

// === track search results
if ( o.is_searchresult )  p.trackSiteSearch( o.search_term, false, o.search_count );

// === track events
if ( window.jQuery ) {
$(document).ready(function(){

	if ( $("#pages.news >.article").length ) { p.trackEvent( 'TZI', 'displayed', 'ARTIKELAUFRUF') };
	if ( $("#pages.news >.article.longread").length ) { p.trackEvent( 'TZI', 'displayed', 'ARTIKELAUFRUF-LONGREAD') };

	if ( $("#pages .body article form").length ) { 
		p.trackEvent( 'Verlag', 'displayed', normalizePath(location.pathname.split(';')[0]) ) ;

		};


	// lookup for whole href because of other url scheme when reload with hints after form submit
	if ( (/172913/).test( location.href ) ) { 

		var formPath = 172913;
		var withHints =  $('.form .sectbody .error').length ? 'withHints ' : 'new ';

		console.log(formPath + ': ' + withHints );
		p.trackEvent( 'Verlag', 'displayed', formPath + ': ' + withHints ) ;

		let checkList = {};
		checkList.keyup = ['#from', '#abo_name','#abo_vorname','#abo_land','#abo_name1','#abo_plz','#abo_ort','#abo_strasse','#abo_tel',
						   '#rg_vorname','#rg_name','#rg_strasse','#rg_land','#rg_plz','#rg_ort','#rg_tel',
						   '#iban','#bic', '#ktoinh', 'textarea[name="bem"]' ];
		checkList.focus = ['#from', '#abo_name','#abo_vorname','#abo_land','#abo_name1','#abo_plz','#abo_ort','#abo_strasse','#abo_tel',
						   '#rg_vorname','#rg_name','#rg_strasse','#rg_land','#rg_plz','#rg_ort','#rg_tel',
						   '#iban','#bic', '#ktoinh', 'textarea[name="bem"]' ];

		checkList.mousedown = [ '#nutzungsbedingungen', '#datenschutz', 'select[name="zahlungsart"]', 'input[type="submit"]' ];

		checkList.keyup.forEach(function(item){
			$( item ).one('keyup ', function(){  
				console.log( formPath + ': ' + item + ' keyup ' + withHints  );
				p.trackEvent( 'Verlag', 'keyup', formPath + ': ' + item + ' keyup ' + withHints    ) ;

			} );
		});
		checkList.focus.forEach(function(item){
			$( item ).one('focus ', function(){  
				console.log( formPath + ': ' + item + ' focus ' + withHints  );
				p.trackEvent( 'Verlag', 'focus', formPath + ': ' + item + ' focus ' + withHints    ) ;  

			} );
		});
		// use this for checkboxes, select options and submit-Button
		checkList.mousedown.forEach(function(item){
			$( item ).one('mousedown ', function(){  
				console.log( formPath + ': ' + item + ' mousedown ' + withHints  );
				p.trackEvent( 'Verlag', 'mousedown', formPath + ': ' + item + ' mousedown ' + withHints    ) ;

			} );
		});

		};

	if ( (/115932/).test( location.href ) ) { 

		var formPath = 115932;
		console.log(formPath  );
		p.trackEvent( 'Verlag', 'displayed', formPath  ) ;

		let checkList = {};
		checkList.keyup = ["input.euro", "input.email", "input.IBAN", "input.fullName" ];
		checkList.focus = ["input.euro", "input.email", "input.IBAN", "input.fullName" ];
		checkList.mousedown = [ "select[name='tzi-praemie[enrol][interval]']", "select[name='tzi-praemie[enrol][start]']",
								"fieldset#tzi-praemie_step:nth(0) button.default.submit", "fieldset#tzi-praemie_step:nth(1) button.default.submit", "fieldset#tzi-praemie_step:nth(2) button.default.submit", 
								"fieldset#tzi-praemie_step:nth(0) button.alt.submit", "fieldset#tzi-praemie_step:nth(1) button.alt.submit", "fieldset#tzi-praemie_step:nth(2) button.alt.submit" ];

		checkList.keyup.forEach(function(item){
			$( document ).one('keyup ', item, function(){  
				console.log( formPath + ': ' + item + ' keyup '  );
				p.trackEvent( 'Verlag', 'keyup', formPath + ': ' + item + ' keyup '   ) ;
			} );
		});
		checkList.focus.forEach(function(item){
			$( document ).one('focus ', item, function(){  
				console.log( formPath + ': ' + item + ' focus '  );
				p.trackEvent( 'Verlag', 'focus', formPath + ': ' + item + ' focus '    ) ;  
			} );
		});
		// use this for checkboxes, select options and submit-Button
		checkList.mousedown.forEach(function(item){
			$( document ).one('mousedown ', item, function(){  
				console.log( formPath + ': ' + item + ' mousedown ' );
				p.trackEvent( 'Verlag', 'mousedown', formPath + ': ' + item + ' mousedown '    ) ;
			} );
		});
	};

	if ( (/170378/).test( location.href ) ) { 

		var formPath = 170378;
		console.log(formPath  );
		p.trackEvent( 'Verlag', 'displayed', formPath  ) ;

		let checkList = {};
		checkList.keyup = ["input.euro", "input.email", "input[name='tzi-aufstocker[Eingabe_Aufstocken][TziNr]']" ];
		checkList.focus = ["input.euro", "input.email", "input[name='tzi-aufstocker[Eingabe_Aufstocken][TziNr]']" ];
		checkList.mousedown = [ "select[name='tzi-aufstocker[Eingabe_Aufstocken][Intervall]']", "select[name='tzi-aufstocker[Eingabe_Aufstocken][Start]']", "input.submit" ];

		checkList.keyup.forEach(function(item){
			$( document ).one('keyup ', item, function(){  
				console.log( formPath + ': ' + item + ' keyup '  );
				p.trackEvent( 'Verlag', 'keyup', formPath + ': ' + item + ' keyup '   ) ;
			} );
		});
		checkList.focus.forEach(function(item){
			$( document ).one('focus ', item, function(){  
				console.log( formPath + ': ' + item + ' focus '  );
				p.trackEvent( 'Verlag', 'focus', formPath + ': ' + item + ' focus '    ) ;  
			} );
		});
		// use this for checkboxes, select options and submit-Button
		checkList.mousedown.forEach(function(item){
			$( document ).one('mousedown ', item, function(){  
				console.log( formPath + ': ' + item + ' mousedown ' );
				p.trackEvent( 'Verlag', 'mousedown', formPath + ': ' + item + ' mousedown '    ) ;
			} );
		});
	};

	if ( $("body.homepage .sect_number-teaser2").length ) {  p.trackEvent( 'TZI', 'displayed', 'COUNTER-HOMEPAGE') };
	if ( $("body.isdir:not(.homepage) #pages.news .sect_number-teaser2").length ) { p.trackEvent( 'TZI', 'displayed', 'COUNTER-RESSORTSEITE') };
	if ( $("#pages.corp .sect_number-teaser2").length ) {  p.trackEvent( 'TZI', 'displayed', 'COUNTER-VERLAGSSEITE') };

	$(document).on("click", "body.homepage .sect_number-teaser2",function(event){ p.trackEvent( 'TZI', 'clicked', 'COUNTER-HOMEPAGE'  ) });
	$(document).on("click", "body.isdir:not(.homepage) #pages.news .sect_number-teaser2",function(event){ p.trackEvent( 'TZI', 'clicked', 'COUNTER-RESSORTSEITE'  ) });
	$(document).on("click", "#pages.corp .sect_number-teaser2",function(event){ p.trackEvent( 'TZI', 'clicked', 'COUNTER-VERLAGSSEITE'  ) });

	if ( $("#pages.news >.article").length && $("#tzi-paywahl-fg").length ) { 
	p.trackEvent( 'TZI', 'displayed', 'ARTIKELAUFRUF_mit_Layer') };
	if ( $("#pages.news >.article").length && !$("#tzi-paywahl-fg").length ) { 
	p.trackEvent( 'TZI', 'displayed', 'ARTIKELAUFRUF_ohne_Layer') };

	$(document).on("click", "#tzi-paywahl-fg .tzi-paywahl__yes"        ,function(event){ p.trackEvent( 'TZI', 'clicked', 'LAYER-JA'            ) });
	$(document).on("click", "#tzi-paywahl-fg .tzi-paywahl__close"      ,function(event){ p.trackEvent( 'TZI', 'clicked', 'LAYER-GERADE-NICHT'  ) });
	$(document).on("click", "#tzi-paywahl-fg .tzi-paywahl__subscriber" ,function(event){ p.trackEvent( 'TZI', 'clicked', 'LAYER-SCHON-DABEI'   ) });
	$(document).on("click", "#tzi-shackle    .tzi-shackle__yes"        ,function(event){ p.trackEvent( 'TZI', 'clicked', 'LASCHE-JA'           ) });

	$(document).on("click", ".sect_end a"                         ,function(event){ p.trackEvent( 'Element', 'clicked', 'ABBINDER'             ) });
	$(document).on("click", "#mainFlyout a"                       ,function(event){ p.trackEvent( 'Element', 'clicked', 'FLYOUT'               ) });
	$(document).on("click", "ul.navbar >li >a"                    ,function(event){ p.trackEvent( 'Element', 'clicked', 'NAVBAR'               ) });

/*  navi-tracking paused          
	$(document).on("click", "ul.navbar.news.newsnavigation a "    ,function(event){ p.trackEvent( 'Navi', 'R_Desktop_MenuImHeader', $(this).find("span").text() ) });
	$(document).on("click", "#mainFlyout .navigation.news a"      ,function(event){ p.trackEvent( 'Navi', 'R_Desktop_Flyout', $(this).find("span").text() )       });
	$(document).on("click", "ul.navbar.corp.newsnavigation a "    ,function(event){ p.trackEvent( 'Navi', 'V_Desktop_MenuImHeader', $(".head >h1 a").text() + ' > ' + $(this).find("span").text() ) });
	$(document).on("click", "#globalnavigation a"                 ,function(event){ p.trackEvent( 'Navi', 'V_Desktop_Schwarze_Navi', $(this).find("span").text() )});
	$(document).on("click", "#mainFlyout .navigation.corp a"      ,function(event){ p.trackEvent( 'Navi', 'V_Desktop_Flyout_tazze', $(this).find("span").text() ) });
	$(document).on("click", "#footer .sitemap > .news a"          ,function(event){ p.trackEvent( 'Navi', 'R_Desktop_Footer', $(this).find("span").text() ) });
	$(document).on("click", "#footer .sitemap > .corp a"          ,function(event){ p.trackEvent( 'Navi', 'V_Desktop_Footer', $(this).find("span").text() ) });
*/ 
	$(document).on("click", "ul.navbar >li.more ul.flyout >li >a" ,function(event){ p.trackEvent( 'Element', 'clicked', 'NAVBARMORE'           ) });
	$(document).on("click", ".insert.dep4534"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUB1'            ) });
	$(document).on("click", ".insert.dep4536"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'TRENNER'              ) });
	$(document).on("click", ".insert.dep4777"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUB2'            ) });
	$(document).on("click", ".insert.dep4219"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUB3'            ) });
	$(document).on("click", ".insert.dep5178"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'MOBEINSCHUB1'         ) });
	$(document).on("click", ".insert.dep5179"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'MOBEINSCHUB2'         ) });
	$(document).on("click", ".insert.dep5180"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'MOBEINSCHUB3'         ) });
	$(document).on("click", ".insert.dep5181"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'MOBEINSCHUB4'         ) });
	$(document).on("click", ".insert.dep4755"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBNORD'         ) });
	$(document).on("click", ".insert.dep4828"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBOEKE'         ) });
	$(document).on("click", ".insert.dep4832"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBPOLITIK'      ) });
	$(document).on("click", ".insert.dep5013"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBSPORT'        ) });
	$(document).on("click", ".insert.dep5020"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBGESELLSCHAFT' ) });
	$(document).on("click", ".insert.dep4956"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBKULTUR'       ) });
	$(document).on("click", ".insert.dep4939"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBPAYWAHL'      ) });
	$(document).on("click", ".insert.dep4704"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'EINSCHUBBERLIN'       ) });
	$(document).on("click", ".sect_seealso a"                     ,function(event){ p.trackEvent( 'Element', 'clicked', 'COMBOLINKBOX'         ) });


// home page all news teaser

// mar: tracking aller Startseiten-Teaser am 04.12. deaktiviert

/* 

	$(document).on("click", ".sect_tdt >ul >li > a" ,function(event){ 
		p.trackEvent( 'Element', 'clicked', 'THEMEN DES TAGES ' + parseInt( $(this).parent().index() + 1 )  ); 
	});

	$(document).on("click", "#pages.news.home >.news.page a.nolead.objlink" ,function(event){ 
		var teaserNr = parseInt(   $(this).parents( 'li.brief' ).index(  'li.brief'  ) + 1  );
		p.trackEvent( 'Element', 'clicked', 'Home > Ressort: Newskamin Miniteaser Nr: ' + teaserNr  ); 
	});
*/
// mar teilreaktivierung Startseitenteaser - nur Newskamin
	$(document).on("click", "#pages.news.home >.news.page.first_page a.article.objlink" ,function(event){ 
		// newskamin or not
		var ressortName = $(this).parents('.news.page').find(' .nose > h2 > a > span ').text() == '' ? 'Newskamin' : $(this).parents('.news.page').find(' .nose > h2 > a > span ').text();
		var teaserNr = ressortName == 'Newskamin' ?  parseInt(   $(this).parents( 'li.article:not(.brief)' ).index(  'li.article:not(.brief)'  ) + 1 ) : parseInt(   $(this).parents( 'li.article' ).index() + 1 );
		var schwerpName = $(this).parents('.sect_spb').find('.secthead > h2 > a > span ').text();
		p.trackEvent( 'Element', 'clicked', 'Home > Ressort: ' +  ressortName + ' ' + schwerpName +  ', Nr: ' + teaserNr  ); 
	});


	$(document).on("click", ".article .tail .sect_related a"  ,function(event){ p.trackEvent( 'Element', 'clicked', 'WEITERE-ARTIKEL' );  });
	$(document).on("click", ".news.article.page .tail .sect_adr a"  ,function(event){p.trackEvent( 'Element', 'clicked', 'Verlagsplatz im Artikel' ) });
	$(document).on("click", ".news.article.page ~ div[class*='thema_clip'] a"  ,function(event){ p.trackEvent( 'Element', 'clicked', 'Schwerpunkt-Clip Artikelebene ' )  });
	$(document).on("click", ".news.page .abtest.testshow a"  ,function(event){ p.trackEvent( 'Element', 'clicked', $( this ).closest(".abtest.testshow" ).attr("class").replace("testshow","").replace("rack","").replace("first_rack","") ) });

	$(document).on("click", ".sect_service a", function(event){
		p.trackEvent( 'Element', 'clicked', 'SERVICEBOX: ' + $(this).attr('href') ) });
	$(document).on("click", ".sectfoot li.shariff-button", function(event){
		p.trackEvent( 'Element', 'clicked', 'SOCIAL-MEDIA-BUTTON: ' + $(this).attr('class').split(" ")[1]  ) 
		});  

	// mar: scroll detection: Add elem to "detectElement" and text to "trackingActionText" arrays
	var detectElement = [ $('.shariff'), $('.article:not(.longread) .sect_related:visible'), $('.news.article.page:not(.longread) .tail .sect_adr:visible'), $('.news.article.page:not(.longread) ~ div[class*="thema_clip"]'), $( 'link + form[action*="172913"] input[type="submit"]' ), $( 'link + form[action*="172913"] input' ).first(), $( 'link + form[action*="172913"] input#iban' ).first(), $(' .body #tzi-praemie '), $(' .body #tzi-aufstocker ')   ];
	var scrollToDetected = []; // detect each only once
	var trackingActionText = [ 'zu Social Media Buttons gescrollt', 'zu weitere Artikel gescrollt', 'zum Verlagsplatz (Artikel) gescrollt', 'zum Schwerpunkt-Clip gescrollt', '172913: Zum Submit-Button gescrollt', '172913: Zum ersten Input-Feld gescrollt', '172913: Zu IBAN gescrollt', '115932: tzi Unterstützer werden Main', '170378: Aufstocker Form Main' ];

	$(window).on('scroll', function() {
	   for (var i = 0; i < detectElement.length; i++) {
		var scrElem = detectElement[i];
		if ( scrollToDetected[i] === undefined && scrElem.length && (elemReached(scrElem)) ) {
			_paq.push([ 'trackEvent', 'Element', 'scrolled', trackingActionText[i] ]);       
			console.log( 'Element', 'scrolled', trackingActionText[i] );       
			scrollToDetected[i] = true;
		}
	  }

	$( ".abtest.testshow" ).each(function( i ) {
		var scrElem = $( this );

		if ( scrollToDetected[i + 1000] === undefined && scrElem.length && (elemReached(scrElem)) ) {
			_paq.push([ 'trackEvent', 'Element', 'scrolled', $( this ).attr('class').replace('testshow','').replace('rack','').replace('first_rack','') ]);       
			scrollToDetected[i + 1000] = true;
		}

		});
	});

	function elemReached(scrElem) {
			var docViewTop = $(window).scrollTop(); var docViewBottom = docViewTop + $(window).height();
			var elemTop = $(scrElem).offset().top + 40; return ((elemTop <= docViewBottom) && (elemTop >= docViewTop))
		}


	function normalizePath( str ){
	  return isNaN(str.slice(-1)) ? str : str + '/' ;
	}

}); // $(document).ready()
} // track events

// === debugging
p.log( 'custom', o );
//console.log( "output _paq: ", _paq );

})(); // piwik customizing

</script><!--IVW--><script> 
var googletag = googletag || {}; googletag.cmd = googletag.cmd || [];

function loadgptScript(){

		$.getScript("https://www.googletagservices.com/tag/js/gpt.js");

(function(){
		var adArea = 'Berlin';      
		var page_id = '5816465';     
		
		
		googletag.cmd.push( function() {
				var pa = googletag.pubads();
				pa.setForceSafeFrame(true);
				pa.collapseEmptyDivs();
				pa.setTargeting('position', [ adArea ]);
//pa.setTargeting('test', ['fireplace']); //       * * * test only ! * * *
//pa.setTargeting('test', ['wallpaper']);   //       * * * test only ! * * *
pa.setTargeting('page_id', ['5816465']);
				if( location.protocol =='https' )
						pa.setTargeting('ssl', ['ja']);
				googletag.enableServices();
				});
		})();
}
// mar: start googletagservices AFTER consent
if ( $.cookie('_sp_enable_dfp_personalized_ads') == 'true' ){
loadgptScript()
} else {
__tcfapi('addEventListener', 2, function(tcData,success){
  if(success && tcData.eventStatus === 'useractioncomplete') {
	if ( $.cookie('_sp_enable_dfp_personalized_ads') === 'true' ){
		loadgptScript()
	  }
  } 
});
}

</script><script>
// expect jQuery already to be loaded
(function() {
if( !window.filOtaz_de )		window.filOtaz_de		= Object;
if( !window.filOtaz_de.ads )		window.filOtaz_de.ads		= Object;
if( !window.filOtaz_de.ads.gpt )	window.filOtaz_de.ads.gpt	= Object;

var sizes = {
	 'artikel_medrec-1'    :[300,250]
	,'artikel_medrec-2'    :[300,250]
	,'artikel_medrec_mitte' :[300,250]
	,'ros_sidebar-1'       :[300,250]
	,'ros_sidebar-2'       :[300,100]
	,'uebersicht_medrec-1' :[300,250]
	,'uebersicht_medrec-2' :[300,250]
	,'uebersicht_medrec-3' :[300,250]
	,'uebersicht_medrec-4' :[300,250]
	,'uebersicht_medrec-5' :[300,250]
	,'uebersicht_board-1'  :[624,150]
	,'ros_billboard-1'     :[870,250]
	,'mobile_banner-1'     :[[300,250],[320,50],[320,100]]
	,'mobile_banner-2'     :[[300,250],[320,50],[320,100]]
	,'mobile_banner-3'     :[[300,250],[320,50],[320,100]]
	,'mobile_banner-4'     :[[300,250],[320,50],[320,100]]
	,'ros_top-links'       :[[120,600],[160,600],[300,600]]
	,'ros_top-mitte'       :[[728,90],[970,90]]
	,'ros_top-rechts'      :[[120,600],[160,600],[300,600]]
	,'ros_leaderboard-1'   :[728,90]
	,'fireplace-mitte'     :[[970,90],[728,90]]
	,'fireplace-links'     :[160,601]
	,'fireplace-rechts'    :[160,600]
	,'fireplace_bg'        :[1,1]
	};
var roadblocks = {};


roadblocks = {
 'ros_top-mitte'   :['fireplace_mitte']
,'ros_top-rechts'  :['fireplace_mitte'] 
};


// general ad handling

// --> to be triggered for ad_zone if ad is (to be) delivered
filOtaz_de.ads.fitIn = function( container ) {	// make space for ad in layout
	var container = $( container );
//console.log('fit '+ container.attr('id') );
	if( !container.is('.ad_zone') ) return;
	var packing = container.parent('.ad_bin');
	packing.addClass('shown');
	container.addClass('ad_zone_shown');
	if( container.hasClass('ad_zone_contentad') ) {
		packing.prev('.sect_leads').addClass('aded_'+( container.height() >120 ?'big' :'small' ) );
		var badge = container.hasClass('ad_zone_sold') ? 'Anzeige' : 'taz-Angebot';
		container.before('<div class="ad_badge">'+ badge +'</div>');
		}
	container.trigger('TAZadInserted').trigger('TAZboxChange');
	};

// --> to be triggered for ad_zone if no ad is (to be) delivered
// redundant?
filOtaz_de.ads.reclaim = function( container ) {	// remove whitespace from non-delivered ad
	var container = $( container );
//console.log('reclaim '+ container.attr('id') );
	if( !container.is('.ad_zone') ) return;
	var packing = container.parent('.ad_bin');
	packing.removeClass('shown');
	container.removeClass('ad_zone_shown');
	if( container.hasClass('ad_zone_contentad') ) {
		packing.prev('.sect_leads').removeClass('aded_big aded_small');
		}
	container.trigger('TAZadRemoved').trigger('TAZboxChange');
	};



var fuzeAdFrame = function( frame ) {	// register frame for messages, fire if one is from our content
console.log('legacy fuzeAdFrame', frame  );
	};


// Hintergrund per Message
var eventMethod = window.addEventListener ? "addEventListener" : "attachEvent";
var messageEvent = eventMethod == "attachEvent" ? "onmessage" : "message";
window[eventMethod](messageEvent, function(e) {
if (e.data.toString().indexOf("color") != -1) {
$("#background").css("background-color", e.data.split(" ")[1]);
}
}, false);


// activate ad frame
var adZoneHandled = {};

var get_id = function( name ) {
return $('div[name=' + name + ']').attr('id');
};

var get_name = function( id ) {
return $('#'+ id ).attr('name');
};

unblock = function( id, rb_id ) {
var name = get_name(id);
var index = roadblocks[name].indexOf(rb_id);
if ( index > -1 ) {
roadblocks[name].splice(index, 1);
}
if ( !roadblocks[name].length ) {
googletag.defineSlot( '/53015287/' + name, sizes[name], id ).addService( googletag.pubads() );
googletag.display( id );
}
};

filOtaz_de.ads.gpt.activateAdFrame = function( id ) {  
	if( !id ) return;
	if( adZoneHandled[id] ) return; adZoneHandled[id] = true; // only one active zone per id

	var packing   = $('#ad_bin_'+ id );
	id = 'ad_zone_'+ id;
	var container = packing.find('>#'+ id );

	var name = container.attr('name');
	if( !name || !sizes[name] ) return;

	// handle custom ad load events
	container.on('filOtaz_de.ads.gpt.noBanner',   function(){  return false;  });
	container.on('filOtaz_de.ads.gpt.preBanner',  function(){
		filOtaz_de.ads.fitIn( container );
		return false;
		});
	container.on('filOtaz_de.ads.gpt.postBanner', function(){  return false;  });


	// define content
	googletag.cmd.push( function() {
		var pa = googletag.pubads();
if ( roadblocks[name] ) {
  pa.addEventListener('slotRenderEnded', function (e) {
	for (var i=0; i<roadblocks[name].length; i++) {
	  if ( e.slot.getSlotElementId() == 'ad_zone_' + roadblocks[name][i] ) {

	if ( e.isEmpty ) {
		  unblock(id, roadblocks[name][i]);
		  }
		else {
		  document.getElementById(id).style.display = 'none';
		  for (var key in roadblocks) {
			if ( !roadblocks.hasOwnProperty(key) ) continue;
			for ( var j=0; j<roadblocks[key].length; j++ )
			if ( id == 'ad_zone_' + roadblocks[key][j] ) {
			  unblock(get_id(key), roadblocks[key][j]);
			  break;
			  }
			}
		  }
		}
	  }
		}); // pa.addEventListener
		//	    console.log('roadblocks' + id);
  googletag.display(id);
  }
else {
  pa.display( '/53015287/'+ name, sizes[name], id );
		// console.log( 'asked for', '/53015287/'+ name, sizes[name], id );
  }
		pa.addEventListener('slotRenderEnded', function(e) {
		if( e.slot.getSlotElementId() != id )  return;	// not our slot
		container.trigger( e.isEmpty ? 'filOtaz_de.ads.gpt.noBanner' : 'filOtaz_de.ads.gpt.preBanner' );
		//console.log( name +( e.isEmpty ? ' isEmpty' : '' ) );
			});
});
}; //filOtaz_de.ads.gpt.activateAdFrame

// activate wallpaper combo add frame
/*
filOtaz_de.ads.gpt.wallpaper = function() {
	console.log('filOtaz_de.ads.gpt.wallpaper not yet supported');
	};
*/

})();
</script><link rel="apple-touch-icon" sizes="57x57" href="/apple-icon-57x57.png"><link rel="apple-touch-icon" sizes="60x60" href="/apple-icon-60x60.png"><link rel="apple-touch-icon" sizes="72x72" href="/apple-icon-72x72.png"><link rel="apple-touch-icon" sizes="76x76" href="/apple-icon-76x76.png"><link rel="apple-touch-icon" sizes="114x114" href="/apple-icon-114x114.png"><link rel="apple-touch-icon" sizes="120x120" href="/apple-icon-120x120.png"><link rel="apple-touch-icon" sizes="144x144" href="/apple-icon-144x144.png"><link rel="apple-touch-icon" sizes="152x152" href="/apple-icon-152x152.png"><link rel="apple-touch-icon" sizes="180x180" href="/apple-icon-180x180.png"><link rel="icon" type="image/png" sizes="192x192" href="/android-icon-192x192.png"><link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png"><link rel="icon" type="image/png" sizes="96x96" href="/favicon-96x96.png"><link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png"><link rel="manifest" href="/manifest.json"><meta name="msapplication-TileColor" content="#ffffff"><meta name="msapplication-TileImage" content="/ms-icon-144x144.png"><link rel="mask-icon" href="/safari-pinned-tab.svg" color="#d50d2e"><meta name="theme-color" content="#ffffff"><script type="text/javascript" src="//data-2d3a3249cd.taz.de/iomb/latest/sensor/manager/base/es6/bundle.js" async="true" crossorigin="anonymous"></script><script type="text/javascript" src="//script.ioam.de/iam.js?m=1" async="true"></script></head><body class="js vga svga xga nosxga"><noscript><img src="https://taz.de/stats/piwik.php?idsite=1&amp;rec=1&amp;action_name=NOSCRIPT" style="border:0" alt=""></img></noscript><script type="text/javascript">(function(){
var body = $(document.body);
body.addClass('js');
if( screen.width >=640 )	body.addClass('vga');	else body.addClass('novga');
if( screen.width >=800 )	body.addClass('svga');	else body.addClass('nosvga');
if( screen.width >=1024 )	body.addClass('xga');	else body.addClass('noxga');
if( screen.width >=1280 )	body.addClass('sxga');	else body.addClass('nosxga');
})();		


	</script><div class="topo_wit" id="counter"><!-- ================ begin: ivw (szm 2.0) ================ --><script xmlns="" type="text/javascript">
IOMm('configure', { 
	st: "taz", 
dn: 'data-2d3a3249cd.taz.de',
mh: 5
	});
</script><script xmlns="" type="text/javascript">
IOMm('pageview', { cp: "Redaktion/Berlin,Artikel" });
</script><!-- ================ end: ivw (szm 2.0) ================ --></div><div id="adOverlay" class="topo_lay"></div><div id="fake" class="topo_wit"></div><div id="background" class="topo_bin"><div id="skirt" style="top: 1500px;"></div><div id="centered"><script type="text/javascript">	// to be called inline early in #centered
//console.log( $('#background') );
//console.log( $('#centered') );
//console.log( $('#pages') );
var checkHash = function() {
	var hashURL = window.location.hash.slice(1);
	if( hashURL.indexOf('!tom=')         ==0 ) new filOtaz_de.TomOL( hashURL.substr(5) );
	if( hashURL.indexOf('!g')            ==0 ) new filOtaz_de.Gallery('/'+ hashURL );
	if( hashURL.indexOf('!vimeo=')       ==0 ) new filOtaz_de.VimeoVideo('/'+ hashURL );
	if( hashURL.indexOf('berlinfolgen')  ==0 ) new patOtaz_de.Berlinfolgen( hashURL );
	if( hashURL.indexOf('!bbi')          ==0 ) new patOtaz_de.BbiFluglaermkarte();
	if( hashURL.indexOf('!pwatch')       ==0 ) new patOtaz_de.ParteispendenWatch();
	if( hashURL.indexOf('!ackerkartell') ==0 ) new patOtaz_de.Ackerkartell();
	if( hashURL.indexOf('!wm2014')       ==0 ) new patOtaz_de.Wm2014();
	if( hashURL.indexOf('!track=')        ==0 ) {
		var tracking_key = hashURL.substr(7);
		console.log( "DEBUG TRACKURL tracking_key="+ tracking_key );
		window.location.hash = '';
		
		if ( tracking_key =="TZISOCIALMEDIA" ) {
			_paq.push([ 'trackEvent', 'Verlag', 'Funnel', 'TZISOCIALMEDIA' ]);
			console.log("DEBUG-piwik TZISOCIALMEDIA");
		}
		if ( tracking_key =="TZIANZEIGE1" ) {
			_paq.push([ 'trackEvent', 'Verlag', 'Funnel', 'TZIANZEIGE1' ]);
			console.log("DEBUG-piwik TZIANZEIGE1");
		}
		if ( tracking_key =="TZIANZEIGE2" ) {
			_paq.push([ 'trackEvent', 'Verlag', 'Funnel', 'TZIANZEIGE2' ]);
			console.log("DEBUG-piwik TZIANZEIGE2");
		}
		if ( tracking_key =="TZIEPAPER" ) {
			_paq.push([ 'trackEvent', 'Verlag', 'Funnel', 'TZIEPAPER' ]);
			console.log("DEBUG-piwik TZIEPAPER");
		}
	}
	//if( hashURL.search(/^![0-9]/) ==0 )	new filOtaz_de.Wtf('/'+ hashURL );
	};
checkHash();
$(window).hashchange( checkHash );
</script><span id="ad_bin_fireplace_bg" class="ad_bin"><div id="ad_zone_fireplace_bg" name="fireplace_bg" class="ad_zone"></div></span><span id="ad_bin_fireplace_mitte" class="ad_bin"><div id="ad_zone_fireplace_mitte" name="fireplace-mitte" class="ad_zone"></div></span><span id="ad_bin_fireplace_links" class="ad_bin"><div id="ad_zone_fireplace_links" name="fireplace-links" class="ad_zone"></div></span><span id="ad_bin_fireplace_rechts" class="ad_bin"><div id="ad_zone_fireplace_rechts" name="fireplace-rechts" class="ad_zone"></div></span><span id="ad_bin_fireplace_script" class="ad_bin"><div id="ad_zone_fireplace_script" name="fireplace-script" class="ad_zone"></div></span><script type="text/javascript">
(function(){
var parent = $('#ad_bin_fireplace_rechts').parent();
if( !parent.is('.wing') || parent.is('.floating.wing') )  // for wings only run when already floating
filOtaz_de.ads.gpt.activateAdFrame('fireplace_rechts');
})(); 
(function(){
var parent = $('#ad_bin_fireplace_bg').parent();
if( !parent.is('.wing') || parent.is('.floating.wing') )  // for wings only run when already floating
filOtaz_de.ads.gpt.activateAdFrame('fireplace_bg');
})(); 
(function(){
var parent = $('#ad_bin_fireplace_mitte').parent();
if( !parent.is('.wing') || parent.is('.floating.wing') )  // for wings only run when already floating
filOtaz_de.ads.gpt.activateAdFrame('fireplace_mitte');
})(); 
(function(){
var parent = $('#ad_bin_fireplace_links').parent();
if( !parent.is('.wing') || parent.is('.floating.wing') )  // for wings only run when already floating
filOtaz_de.ads.gpt.activateAdFrame('fireplace_links');
})(); 
	  </script><div id="adzone_wall" class=""><span id="ad_bin_ros_top_rechts" class="ad_bin"><div id="ad_zone_ros_top_rechts" name="ros_top-rechts" class="ad_zone"></div></span><script type="text/javascript"> (function(){
var domId = 'ros_top_rechts';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script><span id="ad_bin_ros_top_mitte" class="ad_bin"><div id="ad_zone_ros_top_mitte" name="ros_top-mitte" class="ad_zone"></div></span><script type="text/javascript"> (function(){
var domId = 'ros_top_mitte';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script></div><div id="pages" class="news"><ul role="navigation" id="globalnavigation" class="navbar" style="overflow: visible;"><li class="first odd trodd"><a href="https://taz.de/!5811830/" id="menu_p5357"><span>taz Podcast Umfrage</span></a></li><li class="even trodd"><a href="/Abo/!p4209/" id="menu_p4209"><span>Abo</span></a></li><li class="odd treven"><a href="/Genossenschaft/!p4271/" id="menu_p4271"><span>Genossenschaft</span></a></li><li class="even trodd"><a href="https://taz.de/!p4697/#matomo:pk_campaign" id="menu_p5149"><span>taz zahl ich</span></a></li><li class="odd trodd"><a href="/Info/!p4206/" id="menu_p4206"><span>Info</span></a></li><li class="even treven"><a href="/!p4233/" id="menu_p5311"><span>Veranstaltungen</span></a></li><li class="odd trodd"><a href="https://shop.taz.de/#pk_campaign" target="_blank" id="menu_p4378"><span>Shop</span></a></li><li class="more"><a><span>weitere</span></a><ul class="flyout" style="display: none;"><li class="even trodd"><a href="/Anzeigen/!p4288/" id="menu_p4288"><span>Anzeigen</span></a></li><li class="odd treven"><a href="/!p5099/" id="menu_p5106"><span>taz FUTURZWEI</span></a></li><li class="even trodd"><a href="/!p5298/" id="menu_p5319"><span>taz Talk</span></a></li><li class="odd trodd"><a href="https://taz.de/!p4905/" id="menu_p5133"><span>taz lab</span></a></li><li class="even treven"><a href="https://taz.de/!p5122/" id="menu_p5344"><span>taz wird neu</span></a></li><li class="odd trodd"><a href="/!p5297/" id="menu_p4955"><span>taz in der Kritik</span></a></li><li class="even trodd"><a href="/!p4662/" id="menu_p5148"><span>taz am Wochenende</span></a></li><li class="odd treven"><a href="//blogs.taz.de/" id="menu_p4366"><span>Blogs &amp; Hausblog</span></a></li><li class="even trodd"><a href="//monde-diplomatique.de/" id="menu_p4387"><span>LE MONDE diplomatique</span></a></li><li class="odd trodd"><a href="/Thema/!p4786/" id="menu_p4786"><span>Thema</span></a></li><li class="even treven"><a href="/Panter-Stiftung/!p4258/" id="menu_p4258"><span>Panter Stiftung</span></a></li><li class="odd trodd"><a href="/Panter-Preis/!p4207/" id="menu_p4207"><span>Panter Preis</span></a></li><li class="even trodd"><a href="/Recherchefonds-Ausland/!p5062/" id="menu_p5062"><span>Recherchefonds Ausland</span></a></li><li class="odd treven"><a href="/Reisen-in-die-Zivilgesellschaft/!p4310/" id="menu_p4310"><span>Reisen in die Zivilgesellschaft</span></a></li><li class="even trodd"><a href="/!p5044/" id="menu_p5123"><span>Christian Specht</span></a></li><li class="odd trodd"><a href="https://taz.de/!114771/?x" id="menu_p4357"><span>e-Kiosk</span></a></li><li class="even treven"><a href="/Salon/!p5021/" id="menu_p5021"><span>Salon</span></a></li><li class="odd trodd"><a href="/Kantine/!p4237/" id="menu_p4237"><span>Kantine</span></a></li><li class="even trodd"><a href="/Archiv/!p4311/" id="menu_p4311"><span>Archiv</span></a></li></ul></li><li class="last odd treven right"><a href="/Hilfe/!p4591/" id="menu_p4591"><span>Hilfe</span></a></li></ul><div class="full news report article page first odd first_page n1" itemscope="" itemtype="http://schema.org/NewsArticle"><span class="body" role="main" style="min-height: 762px;"><div class="main rack first_rack" id="xid819017">

<div xmlns="" class="metadata" itemscope="">
<a id="articleURL" href="/!5816465"></a><a id="speakingURL" href="/Koalitionsverhandlungen-in-Berlin/!5816465"></a><meta xmlns="http://www.w3.org/1999/xhtml" itemprop="cms-article-ID" content="5816465"><meta xmlns="http://www.w3.org/1999/xhtml" itemprop="url-ID" content="5816465"><meta xmlns="http://www.w3.org/1999/xhtml" itemprop="cms-obj-ID" content="5230729"><meta xmlns="http://www.w3.org/1999/xhtml" itemprop="vgwort-counter-ID" content="59f49821819d4f7d993af4d8b0e9d17a"><span id="articleID"></span><span id="cmsArticleID">5816465</span><img src="//taz.met.vgwort.de/na/59f49821819d4f7d993af4d8b0e9d17a" width="1" height="1" alt=""><img src="/count/vgwort/59f49821819d4f7d993af4d8b0e9d17a.gif" width="1" height="1" alt="">
</div>
<div role="region" id="" class="odd sect sect_article news report"><article class="sectbody" itemprop="articleBody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Koalitionsverhandlungen in Berlin</span><span class="hide">: </span><span>Klimaschutz in die Verfassung</span>
</h1>
<p xmlns="" itemprop="description" class="intro ">Die Neuauflage von Rot-Grün-Rot betont den Klimaschutz, will aber keine neuen Maßstäbe setzen: Ein Zeitplan, wann Berlin klimaneutral sein soll, fehlt.</p>

<a xmlns="" class="full picture" href="/picture/5230729/948/9770710-1.jpeg" target="fullImage" onclick="
var href = this.href;
vHWin=window.open( href , 'fullImage' , 'width=948,height=474' );
vHWin.focus();
return false;
" itemprop="image"><img src="/picture/5230729/624/9770710-1.jpeg" alt="Sitzende Menschen, im Hintergrund Kühlturm" title="Sitzende Menschen, im Hintergrund Kühlturm"></a><p xmlns="" class="caption">Soll ab 2029 nur noch klimaneutral dampfen: Kraftwerk Reuter West <span class="credit"> Foto: dpa</span></p>

<p xmlns="" class="article first odd"><span>BERLIN</span> <em>taz</em> | Mittwochabend auf dem Euref-Campus unter dem Schöneberger Gasometer: Draußen leuchten die Tesla-Schnelladesäulen rot in der Dunkelheit; drinnen, im historischen Wasserturm, läuft die <a target="_blank" href="/Gespraeche-ueber-Rot-Gruen-Rot/!5816336/">siebte Runde der rot-grün-roten Koalitionsverhandlungen</a>. Am Morgen ist die für Energie, Klima-, Umweltschutz und Mobilität zuständige „Dachgruppe“ von demonstrierenden Verkehrswende-AktivistInnen nur bedingt freundlich empfangen worden.</p>
<p xmlns="" class="article even">Als das Spitzenteam Giffey-Jarasch-Lederer deutlich später als angekündigt vor die Presse tritt, stellt es erst einmal klar: Über das Thema Mobilität, das ebenfalls für Mittwoch auf der Tagesordnung stand, wird immer noch geredet. Ergebnisse soll es erst am Freitag geben, wenn auch der Komplex Stadtentwicklung, Bauen und Mieten besprochen wird.</p>

<span id="ad_bin_artikel_medrec_1" class="ad_bin sold contentad"><div id="ad_zone_artikel_medrec_1" name="artikel_medrec-1" class="ad_zone ad_zone_contentad ad_zone_badged ad_zone_sold"></div></span><script type="text/javascript"> (function(){
var domId = 'artikel_medrec_1';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script>

<p xmlns="" class="article odd">Dafür zeigt sich das Trio einig in Sachen Klimaschutz: Man wolle das Thema künftig als Querschnittsaufgabe begreifen, betonen alle drei. Jedes Senatsressort stehe hier mit in der Verantwortung; es werde ein Monitoringsystem entwickelt, um den Fortschritt zu evaluieren und gegebenenfalls nachzusteuern. Wem das zu langweilig klingt, der bekommt ein schickes neues Wort mitgeliefert: „Klima-Governance“.</p>
<p xmlns="" class="article even">Die Neuauflage der Koalition werde den Klimaschutz in die Landesverfassung heben, sagt Klaus Lederer. Er verspricht zudem ein Wärmegesetz, das den Weg aus der Verbrennung fossiler Energieträger vorgeben soll. Mit der Kohle soll es nun schon „2028/2029“ vorbei sein, das nennt der Linken-Chef als neues Verfallsdatum – ab dann soll das Heizkraftwerk Reuter West in Spandau hauptsächlich auf Erdgas umgestellt sein.</p>
<p xmlns="" class="article odd">„Übergangsweise wird man Gas nutzen müssen“, so Lederer, man werde aber mit Nachdruck neue Wärmequellen erschließen, auch Geothermie soll darunter sein. Mit neuen Förderprogrammen soll der Anteil der Stromerzeugung aus Berliner Photovoltaik bis 2035 auf 25 Prozent gesteigert werden.</p>
<p xmlns="" class="article even">Der Pariser 1,5-Grad-Pfad, zu dem sich Jarasch noch einmal bekennt, <a target="_blank" href="/Forscher-ueber-Klimaneutralitaet-2030/!5813510/">ist damit zumindest nach der Ansicht vieler Klimawissenschaftler nicht zu halten</a>, und wohl nicht ohne Grund wollen die VerhandlerInnen keine neue Zeitschiene für das Erreichen der Klimaneutralität definieren.</p>

<span id="ad_bin_artikel_medrec_mitte" class="ad_bin sold contentad"><div id="ad_zone_artikel_medrec_mitte" name="artikel_medrec_mitte" class="ad_zone ad_zone_contentad ad_zone_badged ad_zone_sold"></div></span><script type="text/javascript"> (function(){
var domId = 'artikel_medrec_mitte';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script>

<p xmlns="" class="article odd">R2G hatte erst in diesem Sommer das Ziel der Klimaneutralität von 2050 auf 2045 vorgezogen, ein laufendes Volksbegehren fordert dagegen mit wissenschaftlicher Unterstützung stattdessen das Jahr 2030. „Wir haben mehr über Maßnahmen als über Zielzahlen geredet“, gibt Lederer zu Protokoll, und auch Jarasch meint: „Zielzahlen sind schnell aufgeschrieben.“ Weshalb man es offenbar gerade nicht tun wollte.</p>
<h6 xmlns="">Chefinnensache Gepflegtheit</h6>
<p xmlns="" class="article even">Wer Franziska Giffey zuhört, bekommt ohnehin den Eindruck, das die künftige Regierende lieber „Gepflegtheit“ zur Chefinnensache machen möchte: „Wenn wir über die hehren Ziele des Klimaschutzes sprechen“, so Giffey, dürfe man nicht vergessen, dass auch Sauberkeit eine Stadt lebenswert mache. Die Vermüllung des öffentlichen Raums müsse ein Ende haben: „Die Couch oder die alte Matratze auf der Straße sind nicht gut für das Bild unserer Stadt.“ Man wolle deshalb auch die Straßen und Grünflächenämter der Bezirke besser ausstatten.</p>
<h6 xmlns="">Neue Hoffnung für die Stadtgrün-Charta</h6>
<p xmlns="" class="article last odd">Bei Jarasch klingt das alles etwas moderner und, nun ja, grüner: „Die Stadt soll wieder atmen, wir brauchen mehr Brunnen, mehr Bänke, mehr Bäume.“ Ihr zufolge wird auch ein stagnierendes Projekt schleunigst umgesetzt: „Im Wahlkampf haben wir uns bei diversen Gesetzen verhakt, aber jetzt werden wir die <a target="_blank" href="/Buendnis-fuer-Stadtgruen-Charta/!5812302/">Charta für das Berliner Stadtgrün mit diversen Klarstellungen beschließen</a>.“ Welche das sind – die SPD hatte die Charta wegen ihrer Priorisierung des Freiflächenerhalts blockiert –, auch das wird sich vielleicht am Freitag erhellen.</p>
</article></div>





<div class="sect_text tziBottom"> 									<h6>Ohne Ihre Unterstützung geht es nicht. </h6><p xmlns="" class="article first odd">Da immer mehr Menschen die taz von sich aus unterstützen, können wir auf eine harte Paywall verzichten. 31.500 Menschen sind es schon, die zahlen, nicht weil sie müssen, sondern weil sie es möchten. Weil sie unseren Journalismus, unsere Genauigkeit, unsere kritischen Standpunkte finanzieren und wertschätzen möchten. Sie tun das, weil sie davon überzeugt sind, dass es die taz braucht.  
</p><p xmlns="" class="article even">Das ist unsere Motivation. Klimakatastrophe, Rassismus, soziale Ungleichheit, eine globale Pandemie - gerade jetzt wollen wir mit unseren Artikeln eine relevante Stimme sein. In Zeiten von Desinformation und Verschwörungsmythen ist der freie Zugang zu gut recherchierten Informationen wichtiger denn je. Unsere Artikel auf taz.de sind deshalb frei zugänglich. Wer möchte, kann uns unterstützen. Es gibt aber keinen Bezahlzwang. Diese Idee verfolgt die taz nun seit zehn Jahren, seit 2011 überlassen wir die „Paywahl“ Ihnen, unseren Leser:innen. Schon über 31.500 Menschen haben eine Wahl für die taz getroffen.</p> 										<a target="_blank" href="https://taz.de/!115932/#!formfill:via=Bottom,abTest10,neu" class="button" role="link"><p></p><div class="hint">Voll dabei sein</div></a>								 <div id="tziacc"><div role="region" id="" class="odd sect sect_zahl ich corp"><div class="secthead" role="heading"><h2><a name="zahl ich"><span>zahl ich</span></a></h2><a href="/!163577#Bottom,abTest10,neu;5816465" target="_blank">Einmal zahlen</a></div><div class="sectbody" style="display: none;"><div id="reward" class="reward"><form id="micropay" name="micropay" action="https://taz.de/scripts/taz_zahl_ich/pay_check.php" method="POST">		<input type="submit" disabled="disabled" style="display:none;">		<input type="hidden" size="6" id="ppaid" name="aid" value="taz_online_0">		<input type="hidden" size="6" id="articleid" name="articleid" value="undefined">		<input type="hidden" name="atitle" value="taz.de - online News">		<div><select id="amount" name="amount">				<option value="">Eingabe</option>				<option value="30">0,30 €</option>				<option value="50">0,50 €</option>				<option value="100" selected="true">1,- €</option>				<option value="200">2,- €</option>				<option value="500">5,- €</option>			</select>		<span>oder</span> 	<input type="text" size="6" id="amtman" name="amtman" value=""> 	<span>€</span> 	</div> 	<div>	<input type="hidden" id="pay_way" name="pay_way" value="">			<ul role="toolbar" class="toolbar">				<li><a id="pay_paypal" name="pay_paypal" style="color: rgb(213, 13, 46);" data-original-title="" title="">PayPal</a></li>				<li><a id="pay_ls" name="pay_ls" style="color: rgb(213, 13, 46);" data-original-title="" title="">Lastschrift</a></li>				<li><a id="pay_cc" name="pay_cc" style="color: rgb(213, 13, 46);" data-original-title="" title="">Kreditkarte</a></li>				<li><a id="pay_kto" name="pay_kto" style="color: rgb(213, 13, 46);" data-original-title="" title="">Überweisung</a></li> 			<li><a id="pay_btc" name="pay_btc" style="color: rgb(213, 13, 46);" data-original-title="" title="">Bitcoin</a></li> 		<li class="shariff-button flattr"><a href="https://flattr.com/submit/auto?title=Koalitionsverhandlungen%20in%20Berlin%3A%20Klimaschutz%20in%20die%20Verfassung%20-%20taz.de&amp;description=Die%20Neuauflage%20von%20Rot-Gr%C3%BCn-Rot%20betont%20den%20Klimaschutz%2C%20will%20aber%20keine%20neuen%20Ma%C3%9Fst%C3%A4be%20setzen%3A%20Ein%20Zeitplan%2C%20wann%20Berlin%20klimaneutral%20sein%20soll%2C%20fehlt.&amp;category=text&amp;user_id=null&amp;url=https%3A%2F%2Ftaz.de%2FKoalitionsverhandlungen-in-Berlin%2F!5816465%2F" target="_blank" title="" role="button" aria-label="Artikel flattrn" data-original-title="Artikel flattrn" data-placement="bottom"><span class="fa fa-money"></span><span class="share_text">Flattr</span></a></li></ul>		</div>	</form><ul role="toolbar" class="toolbar"></ul></div><script> filOtaz_de.flattr();</script></div></div></div></div><div id="tzi_abo" style="display:none;" role="region" class="sect sect_tziabo box">    <!-- <div role="heading" class="secthead"><h2>taz zahl ich Abo</h2></div> --></div><div class="tzi-paywahl__logo"><a name="zahl ich" title="Infos über die freiwillige Unterstützung" target="_blank" href="/Infos-ueber-die-freiwillige-Unterstützung/!156925/"></a></div>


<div class="shariff"><ul class="theme-color orientation-horizontal col-8"><li class="shariff-button pocket"><a href="https://getpocket.com/save?url=https%3A%2F%2Ftaz.de%2FKoalitionsverhandlungen-in-Berlin%2F!5816465%2F&amp;title=Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung - taz.de" data-rel="popup" title="" role="button" aria-label="Bei pocket merken" data-original-title="Bei pocket merken" data-placement="bottom"><span class="fa fa-pocket"></span><span class="share_text">merken</span></a></li><li class="shariff-button facebook"><a href="https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Ftaz.de%2FKoalitionsverhandlungen-in-Berlin%2F!5816465%2F" data-rel="popup" title="" role="button" aria-label="Bei Facebook teilen" data-original-title="Bei Facebook teilen" data-placement="bottom"><span class="fa fa-facebook"></span><span class="share_text">teilen</span></a></li><li class="shariff-button twitter"><a href="https://twitter.com/intent/tweet?text=Die%20Neuauflage%20von%20Rot-Gr%C3%BCn-Rot%20betont%20den%20Klimaschutz%2C%20will%20aber%20keine%20neuen%20Ma%C3%9Fst%C3%A4be%20setzen%3A%20Ein%20Zeitplan%2C%20wann%20Berlin%20klimaneutral%20sein%20soll%2C%20fehlt.&amp;url=https%3A%2F%2Ftaz.de%2FKoalitionsverhandlungen-in-Berlin%2F!5816465%2F&amp;via=tazgezwitscher" data-rel="popup" title="" role="button" aria-label="Bei Twitter teilen" data-original-title="Bei Twitter teilen" data-placement="bottom"><span class="fa fa-twitter"></span><span class="share_text">tweet</span></a></li><li class="shariff-button more" data-original-title="" title=""><a aria-label="Mehr Social Media Buttons anzeigen" href="#" role="button" title="" data-original-title="Mehr Social Media Buttons anzeigen" data-placement="bottom"><span class="fa-more_social_buttons fa"></span><span class="share_text"></span></a></li></ul></div>

</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script><script class="tziArticle" type="text/html"><article href="/Headlines2021/!173070/" ratioURL-ressource="173070" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Headlines2021</span><span class="hide">: </span><span>Knapp 32.000 zahlen die Zeche</span>
</h1>
  <p xmlns="" class="article first odd">Viele Menschen zahlen freiwillig, damit taz.de für alle zugänglich bleibt. Es gibt keinen Bezahlzwang, keine Pflicht, keine Paywall, aber auch im Digitalen muss guter Journalismus finanziert werden. Schon beinahe 32.000  Menschen machen mit und finanzieren damit die taz im Netz für alle.       
</p>
  <p xmlns="" class="article even"><strong>Fördern auch Sie jetzt die taz. </strong></p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><script class="tziArticle" type="text/html"><article href="/Headlines2021/!173063/" ratioURL-ressource="173063" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Headlines2021</span><span class="hide">: </span><span>Support your local news dealer </span>
</h1>
  <p xmlns="" class="article first odd">Unsere Artikel sollen alle lesen können – das geht nicht einfach so. Als unabhängiges und frei zugängliches Medium sind wir auf Ihre Unterstützung angewiesen. Nur so können wir unseren Journalismus in  digitalen Zeiten finanzieren und die Arbeit der Redaktion erhalten.        
</p>
  <p xmlns="" class="article even"><strong>Unterstützen Sie jetzt die taz.</strong></p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><script class="tziArticle" type="text/html"><article href="/Headlines2021/!173029/" ratioURL-ressource="173029" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Headlines2021</span><span class="hide">: </span><span>Wir machen Ihnen ein Angebot, das Sie (nicht) ablehnen können.</span>
</h1>
  <p xmlns="" class="article first odd">Jede*r bezahlt, was er*sie möchte. Es gibt keinen Bezahlzwang auf taz.de, aber auch im Digitalen muss die taz finanziert werden. Am  besten freiwillig und solidarisch. Das ist die Idee. Schon über 31.700  Menschen machen mit und finanzieren damit die taz im Netz.        
</p>
  <p xmlns="" class="article even"><strong>Fördern auch Sie jetzt unseren Journalismus. </strong></p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><script class="tziArticle" type="text/html"><article href="/Headlines2021/!172945/" ratioURL-ressource="172945" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Headlines2021</span><span class="hide">: </span><span>Wollen Sie unsere Unterstützer*in sein?</span>
</h1>
  <p xmlns="" class="article first odd">Fast 32.000 machen schon mit und zahlen für unseren Journalismus. Die taz im Netz ist und bleibt damit frei zugänglich – ohne Paywall. Das geht, weil sich unsere Leser*innen an dieser Idee beteiligen. Freiwillig und solidarisch.    
</p>
  <p xmlns="" class="article even"><strong>Unterstützen auch Sie jetzt die taz</strong></p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><script type="text/javascript" xml:space="preserve">
  (function () {
	patOtaz_de.layer_tzi_paywahl({'campaign': '2019-5-27' });
  })();
</script><script class="tziBottomArticle" type="text/html"><article href="/BottomabTest10neu/!155790/" ratioURL-ressource="155790" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Bottom,abTest10,neu</span><span class="hide">: </span><span>Ohne Ihre Unterstützung geht es nicht.</span>
</h1>
  <p xmlns="" class="article first odd">Da immer mehr Menschen die taz von sich aus unterstützen, können wir auf eine harte Paywall verzichten. 31.500 Menschen sind es schon, die zahlen, nicht weil sie müssen, sondern weil sie es möchten. Weil sie unseren Journalismus, unsere Genauigkeit, unsere kritischen Standpunkte finanzieren und wertschätzen möchten. Sie tun das, weil sie davon überzeugt sind, dass es die taz braucht.  
</p>
  <p xmlns="" class="article even">Das ist unsere Motivation. Klimakatastrophe, Rassismus, soziale Ungleichheit, eine globale Pandemie - gerade jetzt wollen wir mit unseren Artikeln eine relevante Stimme sein. In Zeiten von Desinformation und Verschwörungsmythen ist der freie Zugang zu gut recherchierten Informationen wichtiger denn je. Unsere Artikel auf taz.de sind deshalb frei zugänglich. Wer möchte, kann uns unterstützen. Es gibt aber keinen Bezahlzwang. Diese Idee verfolgt die taz nun seit zehn Jahren, seit 2011 überlassen wir die „Paywahl“ Ihnen, unseren Leser:innen. Schon über 31.500 Menschen haben eine Wahl für die taz getroffen.</p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><script class="tziBottomArticle" type="text/html"><article href="/BottomabTest10alt/!155427/" ratioURL-ressource="155427" class="objlink nolead" role="link"><div role="region" id="" class="odd sect sect_article news report"><div class="sectbody"><h1 xmlns="" itemprop="headline">
<span class="kicker">Bottom,abTest10,alt</span><span class="hide">: </span><span>Über 31.500 mal DANKE! Und eine Bitte.</span>
</h1>
  <p xmlns="" class="article first odd">Über 31.500 Personen beteiligen sich bei <em>taz zahl ich.</em> Immer mehr entscheiden sich dafür, uns zu unterstützen. Weil es eine unabhängige, kritische Stimme in den hiesigen Medien braucht. Weil es die taz braucht. Unsere Community ermöglicht damit den freien Zugang für alle. Weil wir Journalismus nicht nur als Produkt oder Ware, sondern auch als öffentliches Gut verstehen. Dies unterscheidet uns von anderen Zeitungen und Bezahlmodellen.                                             
</p>
  <p xmlns="" class="article even">Was uns noch unterscheidet: Unsere Leser:innen. Es sind schon über 31.500, die auf taz.de nichts bezahlen müssten, aber wissen, dass guter Journalismus etwas kostet. Dafür sind wir sehr dankbar – und diesen Schub möchten wir mitnehmen in die Zukunft. Der taz stehen große Veränderungen ins Haus (Stichwort Digitalisierung), die wir nur gemeinsam meistern können. Deshalb suchen wir so viel Unterstützung wie möglich. Es wäre ein schönes Zeichen für die taz und für die Zukunft unseres Journalismus. Mit nur 5,- Euro sind Sie dabei!</p>

  <p xmlns="" class="article last odd"></p>
  </div></div></article></script><div role="region" id="" class="odd sect sect_text "><div class="sectbody"><p>

Fehler auf taz.de entdeckt?
</p><p>
Wir freuen uns über eine 
<a href="mailto:fehlerhinweis@taz.de?subject=Fehlerhinweis%20zu%20https://taz.de/Koalitionsverhandlungen-in-Berlin/!5816465/&amp;body=Zu%20%22Koalitionsverhandlungen in Berlin: Klimaschutz in die Verfassung%22%0D%0Ahttps://taz.de/Koalitionsverhandlungen-in-Berlin/!5816465/%0D%0A%0D%0ABitte%20beschreiben%20Sie%20uns%20den%20Fehler.%20Sie%20können%20gern%20einen%20Screenshot%20anhängen.%20Wir%20danken%20für%20Ihre%20Hilfe%21%20
			%0D%0A%0D%0A
			Inhaltliches%20Feedback%20geben%20Sie%20uns%20gerne%20in%20der%20Kommentarspalte%20oder%20richten%20es%20an%20briefe@taz.de.%20Dort%20leiten%20wir%20ihr%20Feedback%20auch%20an%20die%20Redaktion%20weiter.%20Über%20diese%20Mailadresse%20nehmen%20wir%20nur%20Rechtschreib-%20und%20Faktenkorrekturen%20vor.				">Mail an fehlerhinweis@taz.de</a>!</p><p>
Inhaltliches Feedback? 
</p><p>

Gerne als Leser*innenkommentar unter dem Text auf taz.de oder über das
<a href="/!112355/"> Kontaktformular</a>.
</p></div></div><!-- DEBUG belly: 0 extras here, 0 from directory Berlin, thats 0 merged--></span><script> $(document.getLatestElement()).filter('.wing').on( 'TAZdomChange', function(){
			$(this).find('.rack').removeClass('first first_rack').first().addClass('first first_rack');
			}); //</script><span class="tail" role="complementary"><!-- DEBUG tail: 0 extras here, 0 from directory Berlin, thats 0 merged--><div class="rack first_rack" id="xid479112">

<div role="region" id="" class="first last odd sect sect_profile brief "><div class="secthead" role="heading"><h2><a name="mehr von"><span>mehr von</span></a></h2></div><div class="sectbody" itemscope="itemscope" itemtype="http://schema.org/Person" itemprop="author"><a rel="author" class="brief 





	
	

	author 

community 


		
		
				

	
		



person objlink" itemprop="url" href="/Claudius-Proesser/!a11/"><h4 itemprop="name">Claudius Prößer</h4></a><ul class="contact"><li><a itemprop="email" class="email" href="mailto:Claudius Prößer <claudius@taz.de>"><img src="/lib/ch/web/pix/icons/email.png" class="icon"></a></li></ul></div></div>


</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script><div role="region" id="xid859918" class="even sect sect_adr style_adr "><ul role="directory" class="sectbody corp  directory"><li class=" first last odd trodd online   story leaded leaded pictured"><a class="dept" href="/Mixtape/!p4833/">Mixtape</a><a href="/Freitag-10122021-17-Uhr-ByteFM/!5761606/" ratiourl-ressource="5761606" class="objlink story leaded leaded pictured noavatar" role="link"><h4>Freitag, 10.12.2021, 17 Uhr, ByteFM</h4><h3>taz auf die Ohren</h3><p class="brief">Diesmal im Mixtape: Gewalt, Greg Tate, Parris, Alvin Lucier, FLEE, Spotify goes Militär, Robbie Shakespeare.</p><img class="lozad" src="/picture/4763207/14/byteFM_1280x640turntable.jpg" data-src="/picture/4763207/300/byteFM_1280x640turntable.jpg" alt="" title=""><noscript><img src="/picture/4763207/300/byteFM_1280x640turntable.jpg" alt="" title=""></img></noscript></a></li></ul></div></span><a itemprop="mainEntityOfPage" href="/Koalitionsverhandlungen-in-Berlin/!5816465"></a><div class="float" style="position: absolute; left: 0px; top: 0px; width: 0px; height: 563px;"><div class="rightbar"><span class="wing floating" role="contentinfo" style=""><div class="rack first_rack" id="xid459821">

<!--template mode="layout" match="display[ @obj='descr' ]":start-->


<div role="region" id="" class="even sect sect_meta "><div class="sectbody"><ul role="navigation" class="left rootline toolbar"><li class="first last odd trodd selected"><a href="/Berlin/!p4649/" id="menu_p4649"><span>Berlin</span></a></li></ul><ul role="navigation" class="right toolbar"><li class="date" itemprop="datePublished" content="2021-11-18T11:10:00+01:00">18. 11. 2021</li><li itemprop="publisher" itemscope="" itemtype="http://schema.org/Organization " style=" display:none; "><span itemprop="url" content="http://www.taz.de"></span><span itemprop="name" content="TAZ Verlags- und Vertriebs GmbH"></span><span itemprop="logo" itemscope="" itemtype="http://schema.org/ImageObject"><img itemprop="url" alt="taz logo" src="/lib/ch/web/pix/taz_logo_web.jpg"></span><span content="https://www.facebook.com/taz.kommune/" itemprop="sameAs"></span><span content="https://twitter.com/tazgezwitscher" itemprop="sameAs"></span><span content="https://www.instagram.com/taz.die_tageszeitung" itemprop="sameAs"></span><span content="https://flipboard.com/@taz_de" itemprop="sameAs"></span><span content="https://www.reddit.com/user/dietageszeitung" itemprop="sameAs"></span><span content="https://pod.geraspora.de/u/taz" itemprop="sameAs"></span></li><meta itemprop="dateModified" content="2021-11-18T14:12:39+01:00"></ul></div></div>


<div role="region" id="" class="odd sect sect_profile "><div class="sectbody" itemscope="itemscope" itemtype="http://schema.org/Person" itemprop="author"><a rel="author" class="





	
	

	author 

community 


		
		
				

	
		



person objlink" itemprop="url" href="/Claudius-Proesser/!a11/"><h4 itemprop="name">Claudius Prößer</h4><h5 itemprop="jobTitle">Redakteur taz.Berlin</h5></a><ul class="contact"><li><a itemprop="email" class="email" href="mailto:Claudius Prößer <claudius@taz.de>"><img src="/lib/ch/web/pix/icons/email.png" class="icon"></a></li></ul></div></div>


<div role="region" id="" class="even sect sect_tags "><div class="secthead" role="heading"><h2><a name="Themen"><span>Themen</span></a></h2></div><ul role="directory" class="news directory sectbody"><li class="first odd trodd tag"><a role="link" href="/!t5009370/" class="tag dirlink"><span>Klimaschutz</span></a></li><li class="even trodd tag"><a role="link" href="/!t5364089/" class="tag dirlink"><span>R2G Berlin</span></a></li><li class="odd treven tag"><a role="link" href="/!t5013056/" class="tag dirlink"><span>Koalitionsverhandlungen</span></a></li><li class="last even trodd tag"><a role="link" href="/!t5204208/" class="tag dirlink"><span>Kohleausstieg</span></a></li></ul></div>





<!-- DEBUG label: 0 extras here, 0 from directory Berlin, thats 0 merged-->

</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script><span id="ad_bin_ros_sidebar_1" class="ad_bin sold contentad"><div id="ad_zone_ros_sidebar_1" name="ros_sidebar-1" class="ad_zone ad_zone_contentad ad_zone_badged ad_zone_sold"></div></span><script type="text/javascript"> (function(){
var domId = 'ros_sidebar_1';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script><!-- DEBUG wing: 0 extras here, 0 from directory Berlin, thats 0 merged--></span></div></div></div><div id="tzi-paywahl-bg" style="display: none;"></div><script type="text/javascript">(function () {	// align skirt top with first page's bottom asap
$( document ).ready(function() {
var skirt_top = $('.sect_tdt >.sectbody').length ? $('.sect_tdt >.sectbody').first().offset().top -4 : 1500;
$('#skirt').css({'top': ''+ skirt_top +'px'});
});
})();	</script><div class="full no_rightbar below news page even n2"><div class="head" role="head"><ul role="navigation" class="news navbar newsnavigation"></ul><script>
var liSelected = $(document.getLatestElement()).find('li.selected');
if ( liSelected.length > 1 ){ liSelected.first().removeClass('selected')};
</script></div><span class="body" role="main" style="min-height: 182px;"><div class="rack first_rack" id="xid841997">

<div role="region" id="" class="first last odd sect sect_thema clip_small plain style_clip_small plain "><div class="secthead" role="heading"><h2><a name="Mehr zum Thema"><span>Mehr zum Thema</span></a></h2><ul class="toolbar" role="toolbar"></ul></div><ul role="directory" class="sectbody news  directory"><li class=" first odd trodd    report article leaded pictured"><a href="/Koalitionsverhandlungen-in-Berlin/!5816786/" ratiourl-ressource="5816786" class="objlink report article leaded pictured noavatar" role="link"><h4>Koalitionsverhandlungen in Berlin</h4><h3>Langsam wird es zäh</h3><p>SPD, Grüne und Linke hängen beim Thema Stadtentwicklung und Mietenpolitik. Die Folge: Ein Verhandlungsmarathon kommende Woche.  <span class="author">Bert Schulz</span></p><img class="lozad" src="/picture/5235067/14/0249-2.jpeg" data-src="/picture/5235067/300/0249-2.jpeg" alt="Menschen protestieren für die Umsetzung des Volksentscheids" title="Menschen protestieren für die Umsetzung des Volksentscheids"><noscript><img src="/picture/5235067/300/0249-2.jpeg" alt="Menschen protestieren für die Umsetzung des Volksentscheids" title="Menschen protestieren für die Umsetzung des Volksentscheids"></img></noscript></a></li><li class=" even trodd    report article leaded pictured"><a href="/Koalitionsverhandlungen-in-Berlin/!5816735/" ratiourl-ressource="5816735" class="objlink report article leaded pictured noavatar" role="link"><h4>Koalitionsverhandlungen in Berlin</h4><h3>Verkehrter Ansatz</h3><p>SPD, Grüne und Linke setzen beim Thema Mobiliät vor allem auf Unterstützung vom Bund. Die Verkehrswende sollen die Bezirke umsetzen.  <span class="author">Uwe Rada</span></p><img class="lozad" src="/picture/5234553/14/RGRVerkehr-1.jpeg" data-src="/picture/5234553/300/RGRVerkehr-1.jpeg" alt="Man sieht Lederer, Giffey, Jarasch" title="Man sieht Lederer, Giffey, Jarasch"><noscript><img src="/picture/5234553/300/RGRVerkehr-1.jpeg" alt="Man sieht Lederer, Giffey, Jarasch" title="Man sieht Lederer, Giffey, Jarasch"></img></noscript></a></li><li class=" last odd treven    report article leaded pictured"><a href="/Koalitionsgespraeche-zu-Stadtentwicklung/!5812552/" ratiourl-ressource="5812552" class="objlink report article leaded pictured noavatar" role="link"><h4>Koalitionsgespräche zu Stadtentwicklung</h4><h3>Sei schlau, verlass den Bau</h3><p>SPD, Linke und Grüne streiten um Wohnungsneubau und Regulierung. Ein Bausenator könnte Kompetenzen verlieren. Muss die Linke am Amt festhalten?  <span class="author">Erik Peter, Uwe Rada</span></p><img class="lozad" src="/picture/5232644/14/221507793-1.jpeg" data-src="/picture/5232644/300/221507793-1.jpeg" alt="Sebastain Scheel mit gezücktem Bauhelm auf einer Baustelle" title="Sebastain Scheel mit gezücktem Bauhelm auf einer Baustelle"><noscript><img src="/picture/5232644/300/221507793-1.jpeg" alt="Sebastain Scheel mit gezücktem Bauhelm auf einer Baustelle" title="Sebastain Scheel mit gezücktem Bauhelm auf einer Baustelle"></img></noscript></a></li><a class="button cta">Alle Artikel zum Thema</a></ul></div>

</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script></span><script> $(document.getLatestElement()).filter('.wing').on( 'TAZdomChange', function(){
			$(this).find('.rack').removeClass('first first_rack').first().addClass('first first_rack');
			}); //</script><a itemprop="mainEntityOfPage" href="/Koalitionsverhandlungen-in-Berlin/!5816465"></a><div class="float" style="position: absolute; left: 0px; top: 0px; bottom: 0px; width: 0px; height: 182px;"><div class="rightbar"><div class="nose" role="heading" id="pg_hd" style="z-index: 0;"><h2><a name="Volle Spalte unterm Artikel" href="/Volle-Spalte-unterm-Artikel/!p5315/"><span>Volle Spalte unterm Artikel</span></a></h2></div><span class="headed wing floating" role="complementary" style=""></span></div></div></div><div class="full community page last odd n3"><div class="head" role="head"><ul role="navigation" class=" navbar newsnavigation"></ul><script>
var liSelected = $(document.getLatestElement()).find('li.selected');
if ( liSelected.length > 1 ){ liSelected.first().removeClass('selected')};
</script></div><span class="body" role="main" style="min-height: 570px;"><div class="rack first_rack" id="xid495498">


<div role="region" id="" class="first last odd sect sect_text "><div class="secthead" role="heading"><h2><a name="So können Sie kommentieren:"><span>So können Sie kommentieren:</span></a></h2></div><div class="sectbody"><p>       Bitte <a href="https://www.taz.de/kommune/post.php?a=registration" target="_blank">registrieren Sie sich</a> und halten Sie sich an unsere <a href="http://taz.de/netiquette" target="_blank">Netiquette</a>.     
</p><p>      Haben Sie Probleme beim Kommentieren oder Registrieren?   
</p><p>  Dann mailen Sie uns bitte an <a href="mailto:kommune@taz.de" target="_blank">kommune@taz.de</a>   </p></div></div>

</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script><form name="form" method="post" enctype="application/x-www-form-urlencoded" id="send-comment" class="even sect sect_send-comment_form style_commentlinks js jsForm" action="https://taz.de/kommune/post.php" target="_blank"><div role="heading" class="secthead"><h2><a name="Den ersten Beitrag schreiben"><span>Den ersten Beitrag schreiben</span></a></h2></div><div class="sectbody jsForm"><input class="constant" type="hidden" name="test" value="yes"><input class="constant" type="hidden" name="id" value="5816465"><input type="hidden" name="scope" class="constant" id="form_scope" value="cms-article"><span class="field left narrow preset"><p class="label">Geben Sie Ihren Kommentar hier ein</p><textarea name="post" class="text neutral" id="form_post" maxlength="2000"></textarea></span><span class="commentLength"><span>noch </span><span id="num"></span><span>Zeichen</span></span><input type="submit" class="submit" name="ignore" value="Einloggen und Senden"></div></form><script> filOtaz_de.Form( 'send-comment' ); </script><script type="text/javascript"> (function(){
if( filOtaz_de.ajaxifyForm ) return;

if ('withCredentials' in new XMLHttpRequest() )
var haveCORS = true;
if( typeof XDomainRequest !== "undefined" )
var haveXDR = true;
//console.log('haveCORS: '+ haveCORS );

filOtaz_de.ajaxifyForm = function( id, options ){
/*	options : {
		option : choice, …,
		handlername : function(){…}, …
		}
	options:
		mirrorInput	true/false, default false
	implemented handlers:
		handle200	handle OK Server response instead of default
		on Success	done first in default 200 handler
	to implement:
		handle4xx	handle specific 4xx response from Server
		onReset		done before resetting form to start
		onError		done first in 4xx handlers
*/
var formID = id;
var form1N = $('#'+ id );
if( !form1N.length ) return;
if( form1N[0].submit_via_ajax ) return; form1N[0].submit_via_ajax = true;

var mirror = {  node : options.mirrorInput ? form1N : $()  };
mirror.toShow = mirror.node.find('input, textarea');
mirror.toHide = mirror.node.find('input:submit');

if( !( location.protocol=='https:' || haveCORS ) ) 
	return $('<div class="warning">Um sichere und dennoch bequeme Kommunikation mit unseren Servern zu erlauben, '
			+'sollten sie zu einem aktuelleren Browser wechseln '
			+'oder diese Seite <a href="https:.#'+ formID +'">per HTTPS aufrufen</a></div>'
		);

//		var callback = callback;

var altFormNs = $();	// keep track of form nodes we may insert from server responses
var clear4reset = function(){
	altFormNs.slideUp( function(){
		altFormNs.remove(); altFormNs = $();
		});
	};
var back2start = function( val ){
	if( typeof val == 'string' ) {
		form1N.find('>.sectbody >span >textarea')
			.val( val )
			.trigger('blur')
			;
		}
	// ToDo: show hint -> change filOtaz_de.Form first
	mirror.toShow.prop('disabled', false);
	mirror.toHide.slideDown();
	form1N.slideDown();
	};
var reset = function( val ){
	clear4reset();
	back2start( val );
	};
var handle200 = ( options.handle200 instanceof Function )? options.handle200 : function( data, status, req ){
//console.log( data.documentElement.outerHTML );
	if( options.onSuccess instanceof Function ) options.onSuccess();
	clear4reset();	// hide follow-up forms

	var message = $(data).filter('.message');
	if( message.length && message.text() >'' ) {	// show success message
		var messageN = $('<form id="successMessage" class="sect" />').append( message.addClass('sectbody') );
		$('<input class="submit" type="submit" value="OK" />').appendTo( message ).click( function(){
			messageN.slideUp( function(){  messageN.remove();  });
			back2start('');	// reset to original form
			return false;
			});
		messageN.hide().prependTo( form1N.parent() ).slideDown();
		form1N.slideUp();
		}
	}


var asySubmit = function( formN ){
	var formN = formN;
	var action;
	var clickLock = false;
	formN.find('>.sectbody >input:submit').click( function(){
		if( clickLock ) return false; clickLock = true;
		action = formN.attr('action'); if( !action ) return;	// ToDo: handle missing action
		action = action.replace(/^http:/,'https:');		// ToDo: handle relative action URL

		var handle422 = function( data ) { // we sent incomplete Data. Show forms from response body!
//console.log( data[0].documentElement.outerHTML );
						var formNs = data.find('form');
						if( !formNs.length ) {
							alert('error: server response fails expectations: no form found.');
							return false;
							}

						var anchor = $('<a />').hide()[ mirror.node.length ?'insertAfter' :'insertBefore' ]( formN );

						formN.not( mirror.node ).slideUp( function(){
							formN.not( form1N ).remove();
							});
						mirror.toShow.attr({ disabled : true });
						mirror.toHide.slideUp();

						var last = anchor;
						formNs.each( function(i,e){ var n = $(e);

							var have = $('#'+ n.attr('id') );
							if(! n.children('.secthead').length )
								have.children('.secthead').first().clone(true).prependTo( n );
							var have = have.not( form1N );
							have.slideUp( function(){
								altFormNs = altFormNs.not( have );
								have.remove();					// remove old incarnation
								});

							n.find('script').remove();	// don't import scripts

							n.hide().insertAfter( last );
							anchor.remove();
							last = n;
							altFormNs = altFormNs.add( n );

							/*
							*/
							var ccN = n.find('#captcha-code');
							var captchaCode = ccN.attr('name');
							if( ccN.length && captchaCode )	// this form needs a captcha, make one!
								Recaptcha.create( captchaCode, 'captcha-code', {
									theme: 'red'
									} );
							(new filOtaz_de.Form( n )).cancel( function(){
								 reset();
								});
							asySubmit( n );
							n.slideDown();									
							});
//console.log( altFormNs.length +' forms found' );

			} //handle422


		var post = {};
		formN.find('input,textarea').each( function(i,n){ var N = $(n);
			post[N.attr('name')] = N.val();
			});

		$.ajax({
			type		: 'POST',
			url		: action,
			cache		: false,
			data		: post,
			dataType 	: 'html', // 'xml',
			xhrFields: {
				withCredentials: true
			},

// chrisso: Disabled beforeSend manipulation because first parameter ("req") is now jqXHR and not original XMLHttpRequest object. (see above "xhrFields")
// 
//					beforeSend	: function( req, settings ){
//						req.responseType = 'document';
//						req.withCredentials = true;
//						//req.overrideMimeType('text/xml');
//						},

			success		: function( data, status, req ){	// our post was accepted, show response message!
				handle200( data, status, req );
				},

			error		: function( req, status, message ){

//console.log('status: '+ status +' :: '+ message ); 
				var data = $( req.response );
				if( !data.length ) {	// data is not document
					var text;
					try{	text = req.responseText;
						if( text ) text = text.replace(/http:/g, "");	// have src'es protocol relative
						data = $('<div />').html( text );
//console.log('response text: '+ text );
						} catch(e) {
//console.log('error reading responseText: '+ e );
						}
					}
				if( !data.length ) {
					try{
//console.log( req.response.head.parentElement.outerHTML );
						} catch(e) {
//console.log('error reading response.head: '+ e );
						}
					}



				({	0   : function() {
						alert('Error: Can not see server answer. CORS missing?');	// seems a Cross-Origin problem
//console.log( data );
						},
					200 : function() {	// 200 but error, IE9 ends up here
						handle200( data, status, req );
						},
					422 : function() {  handle422( data );  }
					})[ req.status ]();
				}, //error

			/* not available befor jQ 1.5 
			statusCode : {
				422	: function( req, status, message ){}
				},
			*/

			complete	: function( req, status ){
//try{ console.log('response headers:\n'+ req.getAllResponseHeaders() ); } catch(e){ console.log('getAllResponseHeaders() failed'); }
//console.log('responseXML: '+ req.responseXML );
//if( req.response ) console.log( req.response.head.parentElement.outerHTML )
//else console.log('no data');
				clickLock = false;
				}

			}); //ajax

		return false;
		}); //click
	}; //asySubmit
asySubmit( form1N );
}; //filOtaz_de.ajaxifyForm
})(); </script><script type="text/javascript"> (function(){

var formN = $('#send-comment');	// who i am

if( formN[0].filOtaz_de_handled ) return; formN[0].filOtaz_de_handled = true;	// run only once



// ajaxify myself
filOtaz_de.ajaxifyForm('send-comment', {
mirrorInput	: true ,
onSuccess	: function( data ){
	filOtaz_de.kommune.updateLogin();
//console.log('success handler send-comment '+ $(data).html() );
//			var message = $(data).find('.message');
//			alert( message.length ? message.text() : 'OK' );
	} 
});   


if( formN.parents('.wing').length ) {    // we're in wing
//console.log('send-comment in wing '+ formN.length +', '+ formN.parents('.wing').length +', '+ formN.parents('.float').length );

var multiUse = function() {

// === lend ourselfes to reply-function ===

// shared functions
var onScroll, reset;

// shared data
var form, anchor, prior, bros, dad, uncles, yMin;


$('.community.page >.body >.sect_commentlinks').find('>.sectbody >li, .thread >li').each( function(i,el){	// each comment

	// per comment data
	var refN	= $(el),
		toolbarN	= refN.find('>.toolbar'),
		id,
		onClick
		;

	toolbarN.find('>.reply >a').click( function(){
		// reply specific
		onClick({
			idKey	: 'asid',
			title	: {
				prefix	: '',
				suffix	: ' antworten'
				},
			hint	: 'Antwort bitte hier eingeben',
			submit	: 'antworten'
			});
		} );
	
	toolbarN.find('>.report >a').click( function(){
		// report specific
		onClick({
			a	: 'report',
			idKey	: 'pid',
			title	: {
				prefix  : 'Kommentar von ',
				suffix  : ' melden'
				},
			hint    : 'Problem bitte hier beschreiben',
			submit	: 'melden'
			});
		} );

	toolbarN.css({ visibility:'visible' });


	onClick = function( par ) {

			// determine shared data late
		if( !form ) {
			form = {
				titleN	: formN.find('>.secthead >h2 >a >span'),
				aN	: formN.find('>.sectbody >input[name="a"]'),
				idN	: formN.find('>.sectbody >input[name="asid"]'),
				pidN	: formN.find('>.sectbody >input[name="pid"]'),
				textN	: formN.find('>.sectbody >span >textarea'),
				submitN	: formN.find('>.sectbody >input:submit')
				};
			form.hintN 	= form.textN.prev('p');
			form.org = {
				title	: form.titleN.text(),
				hint	: form.hintN.text(),
				a	: form.aN.attr('value'),
				submit	: form.submitN.attr('value')
				};
			}

		// determine per comment data late
		if( !id )  	id = refN.attr('id').replace(/^[^0-9]*/, '');


		// change presented data to our needs

		if( par.a ) {
			if( !( form.aN  && form.aN.length  ) )  
				form.aN  = $('<input class="constant" type="hidden" />').insertBefore( form.submitN );
			form.aN.attr({	name	: 'a',
					value	: par.a
					});
			}
		else if( form.aN ) {
			form.aN.remove();
			form.aN = null;
			}

		if( par.idKey ) {
			if( form.idN ) form.idN.remove();
			form.idN = $('<input class="constant" type="hidden" />')
				.attr({	name	: par.idKey,
					value	: id
					})
				.insertBefore( form.submitN );
			}
		else if( form.idN ) {
			form.idN.remove();
			form.idN = null;
			}

		form.titleN.empty().append( 
			par.title.prefix, 
			$('<span>'+ refN.find('>.author >h4').text().replace(/[<>\[\]\'"]/g, '') +'</span>').addClass( refN.attr('class') ),
			par.title.suffix
			);

		form.hintN.text(		par.hint	);

		form.submitN.attr({	value :	par.submit	});


		// add reset options
		(new filOtaz_de.Form( formN )).cancel( function(){ reset();	afterLogin.fadeIn(); });
		formN.on('filOtaz_de.Form.done',   reset );

		// geometry
		var	commBot	=   toolbarN.offset().top +   toolbarN.outerHeight(),
			formBot = form.textN.offset().top + form.textN.outerHeight(),
			d	= commBot - formBot
			;
		if( !yMin ) yMin = formBot - $(window).scrollTop();
		if( commBot - $(window).scrollTop() < yMin )  window.scrollTo( 0, commBot - yMin );

		// relocate form to stick w/ referenced comment
		if( !anchor && d >0 ) {
			prior = formN.prev(), bros = formN.nextAll(), dad = formN.parent(), uncles = dad.filter('.rack').nextAll(), afterLogin = $('form#bb-login').nextAll().not('script');
			messageN = $('#successMessage');
			anchor = $('<div class="community wing pin" />')
				.css({  position: 'absolute',
					top     : formN.offset().top +'px',
					left    : dad.offset().left +'px',
					width   : formN.width() +'px'
					})
				.appendTo( document.body )
				;
			if( uncles.length ) {   // we're in rack
				dad.css({  borderBottomStyle: 'none'  });
				anchor.append( $('<div />').addClass( dad.attr('class') ).append( formN, bros ), uncles );
				}
			else	anchor.append( formN, bros );
			}

		if( anchor ) {

			// fading out "meistkommentiert" when answering

			afterLogin.fadeOut();
		// refresh for second ... n time "antworten"
		// commtented out to prevent "neues kommtar" instead of "antworten"
		//	if (formN.hasClass('worked')){
		//		  if( form.aN ) form.aN.remove();
		//		  if( form.idN ) form.idN.remove();
		//	}

			anchor.animate(
				{ 
					top : formN.offset().top + d
				},
			function(){
					$(window).scroll( onScroll );
					} );
				}

		}; //onClick

	} ); //each comment

	onScroll = function() {
		var ref = prior.length ?( prior.offset().top + prior.outerHeight() ) :dad.offset().top;
		// reset clears the whole form, we don't want that onScroll, so we only call relocateForm instead
		// if( formN.offset().top <= ref )  reset();

		// Do not scroll when already commented and there is success message
		if (( formN.offset().top <= ref ) && ! (formN.hasClass('worked')))  relocateForm();
		};
	reset = function() {
		if( reset ) 	formN.off('filOtaz_de.Form.done');
		resetForm();
		relocateForm();
		};
	relocateForm = function() {
		// clean up
		if( onScroll )  $(window).off('scroll.taz_onscroll');

		// relocate form back
		if( anchor ) {
			var 	ref = prior.length ?( prior.offset().top + prior.outerHeight() ) :dad.offset().top,
				d = anchor.offset().top - ref
				;
			anchor.animate({ top: ref }, d, function(){
				dad.append( formN, bros );
				if( uncles.length ) {	// we were in rack
					dad.css({  borderBottomStyle: null  });
					dad.after( uncles );
					}
				if( anchor ) anchor.remove(); anchor = null;
				} );
			}
		};
	resetForm = function() {
		// reset shown data
		if( form.aN      ) form.aN.remove();
		if( form.idN     ) form.idN.remove();
//				if( form.cancelN ) form.cancelN.remove();
		form.titleN.text(		form.org.title	);
		form.hintN.text(		form.org.hint	);
		form.textN.attr({	value : ''		}).trigger('blur');
		form.submitN.attr({	value : form.org.submit	});

		};

}; // multiUse


if( formN.parents('.float').length ) { 	// already relocated into float
	multiUse();
	}
else {					// wait for relocation

	formN.parents('.wing').on('filOtaz_de_float_mounted', function(){  multiUse();  });
	}
} //we're in wing

})(); </script><span id="ad_bin_uebersicht_medrec_4" class="ad_bin sold contentad"><div id="ad_zone_uebersicht_medrec_4" name="uebersicht_medrec-4" class="ad_zone ad_zone_contentad ad_zone_badged ad_zone_sold"></div></span><script type="text/javascript"> (function(){
var domId = 'uebersicht_medrec_4';
var parent = $( '#ad_bin_'+ domId ).parent();
if ( parent.is('.wing') ) {
	filOtaz_de_float.promise.promise().done(function(){		
		filOtaz_de.ads.gpt.activateAdFrame( domId );
	});
} else {
	filOtaz_de.ads.gpt.activateAdFrame( domId );				
}
})(); 
</script></span><script> $(document.getLatestElement()).filter('.wing').on( 'TAZdomChange', function(){
			$(this).find('.rack').removeClass('first first_rack').first().addClass('first first_rack');
			}); //</script><a itemprop="mainEntityOfPage" href="/Koalitionsverhandlungen-in-Berlin/!5816465"></a><div class="float" style="position: absolute; left: 0px; top: 0px; bottom: 0px; width: 0px; height: 582px;"><div class="rightbar"><span class="wing floating" role="complementary" style=""><div class="rack first_rack" id="xid476383">





</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script><form name="form" method="post" enctype="application/x-www-form-urlencoded" id="bb-login" class="even sect sect_bb-login_form style_kommune js jsForm" action="https://taz.de/kommune/post.php" target="_blank"><div role="heading" class="secthead"><ul role="toolbar" class="toolbar"><li class="login"><a href="" target="_blank">Login</a></li><li class="settings"><a href="https://taz.de/kommune/post.php?a=registration" target="_blank">Registrieren</a></li><li class="newpass"><a shape="rect" target="_blank" href="https://portal.taz.de/?a=nopw">Passwort vergessen?</a></li></ul></div><div class="sectbody jsForm" style="display: none;"><input type="hidden" name="a" class="constant" id="form_a" value="login"><span class="field preset"><p class="label">E-Mail</p><input type="text" name="email" class="line neutral" id="form_email"></span><span class="field left narrow preset"><p class="label">Passwort</p><input type="password" name="password" class="secret neutral" id="form_password"></span><input type="submit" class="submit" name="ignore" value="einloggen"></div></form><script> filOtaz_de.Form( 'bb-login' ); </script><script type="text/javascript"> (function(){
if( filOtaz_de.ajaxifyForm ) return;

if ('withCredentials' in new XMLHttpRequest() )
var haveCORS = true;
if( typeof XDomainRequest !== "undefined" )
var haveXDR = true;
//console.log('haveCORS: '+ haveCORS );

filOtaz_de.ajaxifyForm = function( id, options ){
/*	options : {
		option : choice, …,
		handlername : function(){…}, …
		}
	options:
		mirrorInput	true/false, default false
	implemented handlers:
		handle200	handle OK Server response instead of default
		on Success	done first in default 200 handler
	to implement:
		handle4xx	handle specific 4xx response from Server
		onReset		done before resetting form to start
		onError		done first in 4xx handlers
*/
var formID = id;
var form1N = $('#'+ id );
if( !form1N.length ) return;
if( form1N[0].submit_via_ajax ) return; form1N[0].submit_via_ajax = true;

var mirror = {  node : options.mirrorInput ? form1N : $()  };
mirror.toShow = mirror.node.find('input, textarea');
mirror.toHide = mirror.node.find('input:submit');

if( !( location.protocol=='https:' || haveCORS ) ) 
	return $('<div class="warning">Um sichere und dennoch bequeme Kommunikation mit unseren Servern zu erlauben, '
			+'sollten sie zu einem aktuelleren Browser wechseln '
			+'oder diese Seite <a href="https:.#'+ formID +'">per HTTPS aufrufen</a></div>'
		);

//		var callback = callback;

var altFormNs = $();	// keep track of form nodes we may insert from server responses
var clear4reset = function(){
	altFormNs.slideUp( function(){
		altFormNs.remove(); altFormNs = $();
		});
	};
var back2start = function( val ){
	if( typeof val == 'string' ) {
		form1N.find('>.sectbody >span >textarea')
			.val( val )
			.trigger('blur')
			;
		}
	// ToDo: show hint -> change filOtaz_de.Form first
	mirror.toShow.prop('disabled', false);
	mirror.toHide.slideDown();
	form1N.slideDown();
	};
var reset = function( val ){
	clear4reset();
	back2start( val );
	};
var handle200 = ( options.handle200 instanceof Function )? options.handle200 : function( data, status, req ){
//console.log( data.documentElement.outerHTML );
	if( options.onSuccess instanceof Function ) options.onSuccess();
	clear4reset();	// hide follow-up forms

	var message = $(data).filter('.message');
	if( message.length && message.text() >'' ) {	// show success message
		var messageN = $('<form id="successMessage" class="sect" />').append( message.addClass('sectbody') );
		$('<input class="submit" type="submit" value="OK" />').appendTo( message ).click( function(){
			messageN.slideUp( function(){  messageN.remove();  });
			back2start('');	// reset to original form
			return false;
			});
		messageN.hide().prependTo( form1N.parent() ).slideDown();
		form1N.slideUp();
		}
	}


var asySubmit = function( formN ){
	var formN = formN;
	var action;
	var clickLock = false;
	formN.find('>.sectbody >input:submit').click( function(){
		if( clickLock ) return false; clickLock = true;
		action = formN.attr('action'); if( !action ) return;	// ToDo: handle missing action
		action = action.replace(/^http:/,'https:');		// ToDo: handle relative action URL

		var handle422 = function( data ) { // we sent incomplete Data. Show forms from response body!
//console.log( data[0].documentElement.outerHTML );
						var formNs = data.find('form');
						if( !formNs.length ) {
							alert('error: server response fails expectations: no form found.');
							return false;
							}

						var anchor = $('<a />').hide()[ mirror.node.length ?'insertAfter' :'insertBefore' ]( formN );

						formN.not( mirror.node ).slideUp( function(){
							formN.not( form1N ).remove();
							});
						mirror.toShow.attr({ disabled : true });
						mirror.toHide.slideUp();

						var last = anchor;
						formNs.each( function(i,e){ var n = $(e);

							var have = $('#'+ n.attr('id') );
							if(! n.children('.secthead').length )
								have.children('.secthead').first().clone(true).prependTo( n );
							var have = have.not( form1N );
							have.slideUp( function(){
								altFormNs = altFormNs.not( have );
								have.remove();					// remove old incarnation
								});

							n.find('script').remove();	// don't import scripts

							n.hide().insertAfter( last );
							anchor.remove();
							last = n;
							altFormNs = altFormNs.add( n );

							/*
							*/
							var ccN = n.find('#captcha-code');
							var captchaCode = ccN.attr('name');
							if( ccN.length && captchaCode )	// this form needs a captcha, make one!
								Recaptcha.create( captchaCode, 'captcha-code', {
									theme: 'red'
									} );
							(new filOtaz_de.Form( n )).cancel( function(){
								 reset();
								});
							asySubmit( n );
							n.slideDown();									
							});
//console.log( altFormNs.length +' forms found' );

			} //handle422


		var post = {};
		formN.find('input,textarea').each( function(i,n){ var N = $(n);
			post[N.attr('name')] = N.val();
			});

		$.ajax({
			type		: 'POST',
			url		: action,
			cache		: false,
			data		: post,
			dataType 	: 'html', // 'xml',
			xhrFields: {
				withCredentials: true
			},

// chrisso: Disabled beforeSend manipulation because first parameter ("req") is now jqXHR and not original XMLHttpRequest object. (see above "xhrFields")
// 
//					beforeSend	: function( req, settings ){
//						req.responseType = 'document';
//						req.withCredentials = true;
//						//req.overrideMimeType('text/xml');
//						},

			success		: function( data, status, req ){	// our post was accepted, show response message!
				handle200( data, status, req );
				},

			error		: function( req, status, message ){

//console.log('status: '+ status +' :: '+ message ); 
				var data = $( req.response );
				if( !data.length ) {	// data is not document
					var text;
					try{	text = req.responseText;
						if( text ) text = text.replace(/http:/g, "");	// have src'es protocol relative
						data = $('<div />').html( text );
//console.log('response text: '+ text );
						} catch(e) {
//console.log('error reading responseText: '+ e );
						}
					}
				if( !data.length ) {
					try{
//console.log( req.response.head.parentElement.outerHTML );
						} catch(e) {
//console.log('error reading response.head: '+ e );
						}
					}



				({	0   : function() {
						alert('Error: Can not see server answer. CORS missing?');	// seems a Cross-Origin problem
//console.log( data );
						},
					200 : function() {	// 200 but error, IE9 ends up here
						handle200( data, status, req );
						},
					422 : function() {  handle422( data );  }
					})[ req.status ]();
				}, //error

			/* not available befor jQ 1.5 
			statusCode : {
				422	: function( req, status, message ){}
				},
			*/

			complete	: function( req, status ){
//try{ console.log('response headers:\n'+ req.getAllResponseHeaders() ); } catch(e){ console.log('getAllResponseHeaders() failed'); }
//console.log('responseXML: '+ req.responseXML );
//if( req.response ) console.log( req.response.head.parentElement.outerHTML )
//else console.log('no data');
				clickLock = false;
				}

			}); //ajax

		return false;
		}); //click
	}; //asySubmit
asySubmit( form1N );
}; //filOtaz_de.ajaxifyForm
})(); </script><script type="text/javascript"> (function(){

var formID = 'bb-login';
var form = $('#bb-login');

if( !form.length ) return;
if( form[0].filOtaz_de_form_handled ) return; form[0].filOtaz_de_form_handled = true;

var nodes, org, firstrun = true;
if( !filOtaz_de.kommune ) filOtaz_de.kommune = {};
if( !filOtaz_de.kommune.updateLogin ) filOtaz_de.kommune.updateLogin = function(){
	if( !nodes ) nodes = {	// parse DOM lazily
		head	: form.find('>.secthead'),
		body	: form.find('>.sectbody').hide(),
		a	: form.find('input[name="a"]'),
		inputs	: form.find('span').has('input[name="email"],input[name="password"]'),
		submit	: form.find('input.submit'),
		links	: {
			all		: form.find('ul.toolbar'),
			login		: form.find('ul>li.login').click(function(){
						var me = $(this);
//								nodes.body
						me.closest('.secthead').nextAll('.sectbody')	// we might be dynamically inserted, so parse DOM here
						.slideToggle(function(){
							me.toggleClass('active');
							});
						return false;
						}),
			settings	: form.find('ul>li.settings'),
			newpass		: form.find('ul>li.newpass')
			}
		};
	var duration = firstrun ? 0 : $.fx.speeds._default;
	var tazsid = $.cookie('tazsid');
	if(! tazsid ){  // not logged in
			$('#send-comment .submit').val('Einloggen und Senden');
			};

	if( tazsid ) {	// already logged in
			$('#send-comment .submit').val('Abschicken');

		// mar: refresh after login error and then successful login
		if ( $('#bb-login >.sectbody >p.error').length ){


			$('.sect_bb-login').fadeOut().remove();
			$('.sect_bb-login_form').fadeIn();
			};


		var tazid = $.cookie('tazid');
//console.log('	now logged in as '+ tazid );
		if( !org ) org = {	// store original values lazily
			action	: form.attr('action'),
			a	: nodes.a.attr('value'),
			submit	: nodes.submit.attr('value'),
			links	: {
				login		: nodes.links.login,
				settings	: {
					text	: nodes.links.settings.children('a').text()
					},
				newpass		: nodes.links.newpass
				}
			};
		form.attr({ action: 'https://taz.de/kommune/post.php?a=logout' });
//console.log( tazsid );

		nodes.body.slideUp( duration, function(){
			form.addClass('session');
			nodes.a.attr({ value: 'logout' });
			nodes.inputs.hide();
			nodes.submit.attr({ value: 'ausloggen' });
			nodes.body.fadeIn( duration/3 );
			});
		if( tazid ) var profileURL = '/!ku'+ tazid +'/';
		nodes.links.all.fadeTo( duration/2, 0, function(){
			if( nodes.links.login    )  nodes.links.login.detach();
			if( profileURL ) {
				if( !nodes.links.profile ) {
					nodes.links.profile = $('<li><a href="'+ profileURL +'">mein Profil</a></li>');
					nodes.links.settings.before( nodes.links.profile );
					}
				}
			if( nodes.links.settings )  nodes.links.settings.children('a').text('Einstellungen');
			if( nodes.links.newpass  )  nodes.links.newpass.detach();
			nodes.links.all.fadeTo( duration/2, 1 );
			});
		if( profileURL ) {	// tell user their name
			if( !nodes.salute ) {
				nodes.salute = $();
//console.log('go, get profile '+ profileURL );
				$.ajax({
					type		: 'GET',
					url		: profileURL +'c.xml', 
					dataType	: 'xml',
					success		: function( data, status, xhr ){
						if( !nodes.salute ) return;	// we're too late
						var name = $('content >item[type="person"] >name', data ).text();
//console.log('	found '+ name );
						nodes.salute = $('<h2><a>'+ name +'</a></h2>').hide().prependTo( nodes.head );
						nodes.salute.slideDown( duration );
						},
					error		: function( req, status, message ){
//console.log( req.status +' '+ status +' :: '+ message );
						},
					complete	: function( req, status ){
						}
					});
				}
			}
		} //tazsid
	else if( org )	(function(){	// reset logged out state
		var orgwas = org; org = null;
		nodes.body.fadeOut( duration/2, function(){
			nodes.a.attr({		value	: orgwas.a		});
			nodes.inputs.show();
			nodes.submit.attr({	value	: orgwas.submit	});
			form.removeClass('session');
			}); 
		if( nodes.salute )  nodes.salute.slideUp( duration, function(){
			nodes.salute.remove(); nodes.salute = null;
			});
		if( nodes.links.all ) nodes.links.all.fadeTo( duration/2, 0, function(){
			if( nodes.links.profile ) nodes.links.profile.remove(); nodes.links.profile = null;
			nodes.links.settings.before( orgwas.links.login.removeClass('active') );
			nodes.links.settings.children('a').text( orgwas.links.settings.text );
			nodes.links.settings.after(  orgwas.links.newpass );
			nodes.links.all.fadeTo( duration/2, 1 );
			});

		form.attr({		action	: orgwas.action	});
		})();
	firstrun = false;
	};
filOtaz_de.kommune.updateLogin();

var message = filOtaz_de.ajaxifyForm('bb-login', { handle200 : filOtaz_de.kommune.updateLogin } );

if( message && message.length )	// secure ajax not supported
	form.after( message );

})();
</script><div class="rack" id="xid524032">


<div role="region" id="" class="first last odd sect sect_shop style_shop "><div class="secthead" role="heading"><h2><a name="meistkommentiert"><span>meistkommentiert</span></a></h2><ul class="toolbar" role="toolbar"></ul></div><ul role="directory" class="sectbody   directory"><li class=" first odd trodd    leaded picture-missing"><a href="/Scholz-eroeffnet-naechsten-Wahlkampf/!5818663/" class="objlink brief leaded noavatar" role="link"><h4>Scholz eröffnet nächsten Wahlkampf</h4><h3>Die Kalküle des Kanzlers</h3><p>Die SPD setzt auf Respekt und „normale“ Menschen – in Abgrenzung gegen Eliten und den „woken“ Mainstream. Das ist ein gefährliches Spiel.  </p></a></li><li class=" even trodd   "><a href="/Impfpflicht-in-Gesundheitsberufen/!5821531/" class="objlink noavatar" role="link"><h4>Impfpflicht in Gesundheitsberufen</h4><h3>Eine Art Berufsverbot</h3></a></li><li class=" odd treven   "><a href="/Die-Ampel-Politik-wird-weiblicher/!5818635/" class="objlink noavatar" role="link"><h4>Die Ampel-Politik wird weiblicher</h4><h3>Sicherheit ist Frauensache</h3></a></li><li class=" even trodd   "><a href="/Kevin-Kuehnert-ueber-sein-neues-Amt/!5818642/" class="objlink noavatar" role="link"><h4>Kevin Kühnert über sein neues Amt</h4><h3>„Ohne ein Arschloch zu werden“</h3></a></li><li class=" last odd trodd   "><a href="/Rituale-der-Machtuebergabe/!5818662/" class="objlink noavatar" role="link"><h4>Rituale der Machtübergabe</h4><h3>Als Olaf Scholz sitzen blieb</h3></a></li></ul></div>


</div><script> $(document.getLatestElement()).filter('.rack').trigger('TAZdomChange'); //
	</script></span></div></div></div><div id="footer"><ul role="navigation" class="sitemap"><li class="news"><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a itemprop="url" href="/"><span itemprop="title">taz</span></a></span><ul><li class="first odd trodd"><a href="/Politik/!p4615/" id="menu_p4615"><span>Politik</span></a><ul><li class="first odd trodd"><a href="/Politik/Deutschland/!p4616/" id="menu_p4616"><span>Deutschland</span></a></li><li class="even trodd"><a href="/Politik/Europa/!p4617/" id="menu_p4617"><span>Europa</span></a></li><li class="odd treven"><a href="/Politik/Amerika/!p4618/" id="menu_p4618"><span>Amerika</span></a></li><li class="even trodd"><a href="/Politik/Afrika/!p4621/" id="menu_p4621"><span>Afrika</span></a></li><li class="odd trodd"><a href="/Politik/Asien/!p4619/" id="menu_p4619"><span>Asien</span></a></li><li class="even treven"><a href="/Politik/Nahost/!p4620/" id="menu_p4620"><span>Nahost</span></a></li><li class="last odd trodd"><a href="/Politik/Netzpolitik/!p4622/" id="menu_p4622"><span>Netzpolitik</span></a></li></ul></li><li class="even trodd"><a href="/Oeko/!p4610/" id="menu_p4610"><span>Öko</span></a><ul><li class="first odd trodd"><a href="/Oeko/Oekonomie/!p4623/" id="menu_p4623"><span>Ökonomie</span></a></li><li class="even trodd"><a href="/Oeko/Oekologie/!p4624/" id="menu_p4624"><span>Ökologie</span></a></li><li class="odd treven"><a href="/Oeko/Arbeit/!p4629/" id="menu_p4629"><span>Arbeit</span></a></li><li class="even trodd"><a href="/Oeko/Konsum/!p4625/" id="menu_p4625"><span>Konsum</span></a></li><li class="odd trodd"><a href="/Oeko/Verkehr/!p4628/" id="menu_p4628"><span>Verkehr</span></a></li><li class="even treven"><a href="/Oeko/Wissenschaft/!p4636/" id="menu_p4636"><span>Wissenschaft</span></a></li><li class="last odd trodd"><a href="/Oeko/Netzoekonomie/!p4627/" id="menu_p4627"><span>Netzökonomie</span></a></li></ul></li><li class="odd treven"><a href="/Gesellschaft/!p4611/" id="menu_p4611"><span>Gesellschaft</span></a><ul><li class="first odd trodd"><a href="/Gesellschaft/Alltag/!p4632/" id="menu_p4632"><span>Alltag</span></a></li><li class="even trodd"><a href="/Gesellschaft/Reportage-und-Recherche/!p5265/" id="menu_p5265"><span>Reportage und Recherche</span></a></li><li class="odd treven"><a href="/Gesellschaft/Debatte/!p4633/" id="menu_p4633"><span>Debatte</span></a></li><li class="even trodd"><a href="/Gesellschaft/Kolumnen/!p4634/" id="menu_p4634"><span>Kolumnen</span></a></li><li class="odd trodd"><a href="/Gesellschaft/Medien/!p4630/" id="menu_p4630"><span>Medien</span></a></li><li class="even treven"><a href="/Gesellschaft/Bildung/!p4635/" id="menu_p4635"><span>Bildung</span></a></li><li class="odd trodd"><a href="/Gesellschaft/Gesundheit/!p4637/" id="menu_p4637"><span>Gesundheit</span></a></li><li class="even trodd"><a href="/Gesellschaft/Reise/!p4638/" id="menu_p4638"><span>Reise</span></a></li><li class="last odd treven"><a href="/!p5334/" id="menu_p5318"><span>Podcasts</span></a></li></ul></li><li class="even trodd"><a href="/Kultur/!p4639/" id="menu_p4639"><span>Kultur</span></a><ul><li class="first odd trodd"><a href="/Kultur/Musik/!p4640/" id="menu_p4640"><span>Musik</span></a></li><li class="even trodd"><a href="/Kultur/Film/!p4641/" id="menu_p4641"><span>Film</span></a></li><li class="odd treven"><a href="/Kultur/Kuenste/!p4642/" id="menu_p4642"><span>Künste</span></a></li><li class="even trodd"><a href="/Kultur/Buch/!p4643/" id="menu_p4643"><span>Buch</span></a></li><li class="last odd trodd"><a href="/Kultur/Netzkultur/!p4631/" id="menu_p4631"><span>Netzkultur</span></a></li></ul></li><li class="odd trodd"><a href="/Sport/!p4646/" id="menu_p4646"><span>Sport</span></a><ul><li class="first odd trodd"><a href="/Sport/Fussball/!p4647/" id="menu_p4647"><span>Fußball</span></a></li><li class="last even trodd"><a href="/Sport/Kolumnen/!p4648/" id="menu_p4648"><span>Kolumnen</span></a></li></ul></li><li class="even treven selected" itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a itemprop="url" href="/Berlin/!p4649/" id="menu_p4649"><span itemprop="title">Berlin</span></a><ul></ul></li><li class="odd trodd"><a href="/Nord/!p4650/" id="menu_p4650"><span>Nord</span></a><ul><li class="first odd trodd"><a href="/Nord/Hamburg/!p4651/" id="menu_p4651"><span>Hamburg</span></a></li><li class="even trodd"><a href="/Nord/Bremen/!p4652/" id="menu_p4652"><span>Bremen</span></a></li><li class="last odd treven"><a href="/Nord/Kultur/!p4653/" id="menu_p4653"><span>Kultur</span></a></li></ul></li><li class="last even trodd"><a href="/Wahrheit/!p4644/" id="menu_p4644"><span>Wahrheit</span></a><ul><li class="first odd trodd"><a href="/bei-Tom/!t5180734/" id="menu_p4685"><span>bei Tom</span></a></li><li class="last even trodd"><a href="/ueber-die-Wahrheit/!5068762/" id="menu_p4684"><span>über die Wahrheit</span></a></li></ul></li></ul></li><li class="corp"><ul><li class="first odd trodd"><a href="https://taz.de/!5811830/" id="menu_p5357"><span>taz Podcast Umfrage</span></a></li><li class="even trodd"><a href="/Abo/!p4209/" id="menu_p4209"><span>Abo</span></a></li><li class="odd treven"><a href="/Genossenschaft/!p4271/" id="menu_p4271"><span>Genossenschaft</span></a></li><li class="even trodd"><a href="https://taz.de/!p4697/#matomo:pk_campaign" id="menu_p5149"><span>taz zahl ich</span></a></li><li class="odd trodd"><a href="/Info/!p4206/" id="menu_p4206"><span>Info</span></a></li><li class="even treven"><a href="/!p4233/" id="menu_p5311"><span>Veranstaltungen</span></a></li><li class="odd trodd"><a href="https://shop.taz.de/#pk_campaign" target="_blank" id="menu_p4378"><span>Shop</span></a></li><li class="even trodd"><a href="/Anzeigen/!p4288/" id="menu_p4288"><span>Anzeigen</span></a></li><li class="odd treven"><a href="/!p5099/" id="menu_p5106"><span>taz FUTURZWEI</span></a></li><li class="even trodd"><a href="/!p5298/" id="menu_p5319"><span>taz Talk</span></a></li><li class="odd trodd"><a href="https://taz.de/!p4905/" id="menu_p5133"><span>taz lab</span></a></li><li class="even treven"><a href="https://taz.de/!p5122/" id="menu_p5344"><span>taz wird neu</span></a></li><li class="odd trodd"><a href="/!p5297/" id="menu_p4955"><span>taz in der Kritik</span></a></li><li class="even trodd"><a href="/!p4662/" id="menu_p5148"><span>taz am Wochenende</span></a></li><li class="odd treven"><a href="//blogs.taz.de/" id="menu_p4366"><span>Blogs &amp; Hausblog</span></a></li><li class="even trodd"><a href="//monde-diplomatique.de/" id="menu_p4387"><span>LE MONDE diplomatique</span></a></li><li class="odd trodd"><a href="/Thema/!p4786/" id="menu_p4786"><span>Thema</span></a></li><li class="even treven"><a href="/Panter-Stiftung/!p4258/" id="menu_p4258"><span>Panter Stiftung</span></a></li><li class="odd trodd"><a href="/Panter-Preis/!p4207/" id="menu_p4207"><span>Panter Preis</span></a></li><li class="even trodd"><a href="/Recherchefonds-Ausland/!p5062/" id="menu_p5062"><span>Recherchefonds Ausland</span></a></li><li class="odd treven"><a href="/Reisen-in-die-Zivilgesellschaft/!p4310/" id="menu_p4310"><span>Reisen in die Zivilgesellschaft</span></a></li><li class="even trodd"><a href="/!p5044/" id="menu_p5123"><span>Christian Specht</span></a></li><li class="odd trodd"><a href="https://taz.de/!114771/?x" id="menu_p4357"><span>e-Kiosk</span></a></li><li class="even treven"><a href="/Salon/!p5021/" id="menu_p5021"><span>Salon</span></a></li><li class="odd trodd"><a href="/Kantine/!p4237/" id="menu_p4237"><span>Kantine</span></a></li><li class="even trodd"><a href="/Archiv/!p4311/" id="menu_p4311"><span>Archiv</span></a></li><li class="last odd treven"><a href="/Hilfe/!p4591/" id="menu_p4591"><span>Hilfe</span></a></li></ul></li><li class="help"><ul id="legal"><li class="first odd trodd"><a href="/Hilfe/!p4591/" id="menu_p4591"><span>Hilfe</span></a></li><li class="even trodd"><a href="https://taz.de/!112354/" id="menu_p4679"><span>Impressum</span></a></li><li class="odd treven"><a href="http://taz.de//!173033/" id="menu_p4800"><span>Service</span></a></li><li class="even trodd"><a href="https://taz.de/!114802/" id="menu_p4965"><span>Redaktionsstatut</span></a></li><li class="odd trodd"><a href="https://taz.de/!114191/" id="menu_p4680"><span>RSS</span></a></li><li class="even treven"><a href="https://taz.de/!166598/" id="menu_p4681"><span>Datenschutz</span></a></li><li class="odd trodd"><a href="https://taz.de/!p4283/" id="menu_p4827"><span>Newsletter</span></a></li><li class="even trodd"><a href="https://taz.de/!p4858/" id="menu_p4859"><span>Informant</span></a></li><li class="last odd treven"><a href="https://taz.de/!112355/" id="menu_p4682"><span>Kontakt</span></a></li></ul></li><li class="search"><form class="search" role="search" action="/!s=/" method="POST"><div class="frame"><input type="text" minlength="2" required="required" placeholder="suchen ..." class="text preset" name="s"><input type="image" class="send" title="suche senden" alt="suche senden" name="ignore" src="/lib/ch/web/pix/redArrowsRight_d50d2e.png"></div></form></li></ul><div class="copyright">Alle Rechte vorbehalten. Für Fragen zu Rechten oder Genehmigungen wenden Sie sich bitte an lizenzen@taz.de<a id="mobileswitcher" href="#" onclick="
			event.preventDefault();
			var currentDate = new Date();
			var expDate = new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate()+1, 3, 59, 59);
			$.cookie.raw = true;
			$.cookie('ratioURL_channel', 'moby/force', { expires: expDate, domain:'taz.de', path:'/', secure: false } ); 

			if( (/moby/).test($.cookie('ratioURL_channel') ) ) {

				location.search = ( location.search ? location.search +'&amp;' : '?' ) +'nocache='+ (new Date()).getTime();
				}
			else alert('Cookie konnte nicht gesetzt werden');
			
		"><span>zur mobilen Ansicht wechseln</span></a></div></div><div class="float" style="position: absolute; left: 0px; top: 0px; width: 0px; height: 716px;"><div class="head" role="head"><ul role="navigation" class="news navbar newsnavigation"><li class="first odd trodd home"><a href="/" id="menu_p4608"><span>taz</span><div class="tazze logo"></div></a></li><li class="even trodd"><a href="/Politik/!p4615/" id="menu_p4615"><span>Politik</span></a></li><li class="odd treven"><a href="/Oeko/!p4610/" id="menu_p4610"><span>Öko</span></a></li><li class="even trodd"><a href="/Gesellschaft/!p4611/" id="menu_p4611"><span>Gesellschaft</span></a></li><li class="odd trodd"><a href="/Kultur/!p4639/" id="menu_p4639"><span>Kultur</span></a></li><li class="even treven"><a href="/Sport/!p4646/" id="menu_p4646"><span>Sport</span></a></li><li class="odd trodd selected"><a href="/Berlin/!p4649/" id="menu_p4649"><span>Berlin</span></a></li><li class="even trodd"><a href="/Nord/!p4650/" id="menu_p4650"><span>Nord</span></a></li><li class="last odd treven"><a href="/Wahrheit/!p4644/" id="menu_p4644"><span>Wahrheit</span></a></li></ul><script>
var liSelected = $(document.getLatestElement()).find('li.selected');
if ( liSelected.length > 1 ){ liSelected.first().removeClass('selected')};
</script><form class="search" role="search" action="/!s=/" method="POST"><div class="frame"><input type="text" minlength="2" required="required" placeholder="suchen ..." class="text preset" name="s"><input type="image" class="send" title="suche senden" alt="suche senden" name="ignore" src="/lib/ch/web/pix/redArrowsRight_d50d2e.png"></div></form><div class="pin"><div id="mainFlyout" class="flyout" style="display: none; height: 0px;"></div></div></div></div></div></div></div><iframe name="__tcfapiLocator" style="display: none;"></iframe><div class="tziFgContainer2021 minimized closed thx" id="tzi-paywahl-fg" style="height: 44px; top: auto; bottom: 0px;"><div class="inner"><ul class="article inner"><li class="tzi-paywahl__close"><a title="Wir handeln aus der Überzeugung heraus, dass Worte die Welt verändern können. Deshalb machen wir unsere Inhalte auf taz.de kostenfrei zugänglich. Gerade in Zeiten von Falschmeldungen, Bezahlschranken und Rechtspopulismus müssen Menschen sich unabhängig informieren können. Wenn auch Sie für einen freien, kritischen Journalismus im Netz einstehen wollen, unterstützen Sie uns: Werden Sie Pressefreiheitskämpfer*in!">Gerade nicht</a></li></ul><h1 xmlns="" itemprop="headline">
<span class="kicker">Headlines2021</span><span class="hide">: </span><span>Support your local news dealer </span>
</h1><p xmlns="" class="article first odd">Unsere Artikel sollen alle lesen können – das geht nicht einfach so. Als unabhängiges und frei zugängliches Medium sind wir auf Ihre Unterstützung angewiesen. Nur so können wir unseren Journalismus in  digitalen Zeiten finanzieren und die Arbeit der Redaktion erhalten.        
</p><p xmlns="" class="article even"><strong>Unterstützen Sie jetzt die taz.</strong></p><ul class="article"><li class="tzi-paywahl__yes"><a class="hint" title="Toll, dass Sie unabhängigen Journalismus möglich machen." tabindex="3" target="_blank" href="https://taz.de/!115932/#!formfill:via=Layer,2021&amp;referrer=/Koalitionsverhandlungen-in-Berlin/!5816465/"><span class="tzi-shackle__yes">Ja, ich will</span></a></li><li class="minimized__text"><p class="article"><strong>Unterstützen Sie die taz jetzt freiwillig mit Ihrem Beitrag</strong></p></li><li class="thx__text"><p class="article"><strong>Vielen Dank, dass Sie die taz unterstützen</strong></p></li><li class="tzi-paywahl__close subscriber"><a title="Nicht mehr anzeigen"></a></li><li class="tzi-paywahl__subscriber"><a title="Danke!">Schon dabei!</a></li><li class="tzi-paywahl__logo"><a name="zahl ich" title="Infos über die freiwillige Unterstützung" target="_blank" href="/Infos-ueber-die-freiwillige-Unterstützung/!156925/"></a></li></ul></div></div><script type="text/javascript">
	const lazyObserver = lozad('.lozad', {
		rootMargin: '1050px', // syntax similar to that of CSS Margin
		threshold: 0.1, // ratio of element convergence
		enableAutoReload: true // it will reload the new image when validating attributes changes
	});
	lazyObserver.observe();

</script>
</body><!-- DEBUG end 21:53:31+01:00--></html>`
