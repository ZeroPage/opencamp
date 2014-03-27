var height = viewport().height;

var $intro = document.getElementById("Intro");
$intro.style.height = (height * 2) + "px";

var s = skrollr.init();

function viewport() {
	var e = window , a = 'inner';
	if (!('innerWidth' in window)) {
		a = 'client';
		e = document.documentElement || document.body;
	}
	return { width : e[a + 'Width'] , height : e[a + 'Height'] }
}

