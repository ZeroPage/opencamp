(function(){
	document.addEventListener("touchmove", handleScroll, false);
	document.addEventListener("scroll", handleScroll, false);
	document.addEventListener("hashchanged", handleScroll, false);

	var elements = document.querySelectorAll("[data-menu]");
	elements = [].map.call(elements, function(element){
		var targetId = element.getAttribute("data-menu");
		var target = document.getElementById(targetId);
		return {
			element : element,
			target : target,
			pos : getPositionOfElement(target),
			hash : element.getAttribute("data-menu")
		};
	}).sort(function(a, b){
		return a.pos > b.pos;
	});

	function handleScroll(){
		//var currentPos = document.documentElement.scrollTop || document.body.scrollTop;
		var currentPos = $(document).scrollTop();
		var $target = null;

		for(var i = elements.length - 1; i >= 0; i--){
			if(!$target && elements[i].pos < currentPos + 50){ //must sync with stylesheet
				$target = elements[i];
			}
			elements[i].element.classList.remove("active");
		}
		
		if($(window).scrollTop() + $(window).height() > $(document).height() - 100) {
			$target = elements[elements.length-1];
		}

		if($target) {
			$target.element.classList.add("active");
		}
	}
	function getPositionOfElement(domElement) {
		var pos = 0;
		while (domElement != null) {
			pos += domElement.offsetTop;
			domElement = domElement.offsetParent;
		}
		return pos;
	}
})();

