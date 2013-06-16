//Microsoft.Glimmer.OneWay
//<AnimationCollection FilePath="G:\myprojects\tribe2\glimmer\Glimmer\glimmerUI\glimmerUI\samples\samples\js\freestyle.html.glimmer.js" xmlns="clr-namespace:GlimmerLib;assembly=GlimmerLib" xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"><Animation Name="planeFly" EventType="load" Trigger="{x:Null}"><Animation.Targets><Target Name="#plane" Duration="5000" Easing="easeInElastic" Callback="null"><Target.Effects><YTranslationEffect CSSName="top" DisplayName="Y Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="-180" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /><XTranslationEffect CSSName="left" DisplayName="X Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="820" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="planedropTimer" EventType="load" Trigger="{x:Null}"><Animation.Targets><Target Name="#plane" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="planeDrop" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="5000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="planeDrop" EventType="load" Trigger="{x:Null}"><Animation.Targets><Target Name="#plane" Duration="300" Easing="easeOutBounce" Callback="null"><Target.Effects><YTranslationEffect CSSName="top" DisplayName="Y Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="0" IsStartValue="False" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="cloudsMove" EventType="mouseover" Trigger="#clouds"><Animation.Targets><Target Name="#clouds" Duration="1000" Easing="linear" Callback="null"><Target.Effects><EffectCollection /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="frogIn" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#frog" Duration="300" Easing="swing" Callback="null"><Target.Effects><TimerEffect CSSName="conversation" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="1000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><YTranslationEffect CSSName="top" DisplayName="Y Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="5" IsStartValue="False" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /><XTranslationEffect CSSName="left" DisplayName="X Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="15" IsStartValue="False" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="frogInTimer" EventType="load" Trigger="{x:Null}"><Animation.Targets><Target Name="" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="frogIn" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="6500" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="conversation" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="howdyOut" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="2000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="howdyOut" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="okIn" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="1000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="hidden" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="okIn" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble2" Duration="1000" Easing="easeInBack" Callback="null"><Target.Effects><TimerEffect CSSName="okOut" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="2000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="okOut" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble2" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="sureIn" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="1000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="hidden" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="sureIn" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="mushroom" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="2000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="background-image" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="url('sure.png')" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="mushroom" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#bubble" Duration="1000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="logoIn" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="3000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="background-image" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="url('mushroom.png')" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="logoIn" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#mixLogo" Duration="1000" Easing="easeOutElastic" Callback="null"><Target.Effects><TimerEffect CSSName="textIn" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="2000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><XTranslationEffect CSSName="left" DisplayName="X Position Animation" MaxValue="5000" MinValue="-5000" From="1400px" To="0" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="littlePlane" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#plane2" Duration="12000" Easing="linear" Callback="null"><Target.Effects><TimerEffect CSSName="download" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="2000" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><XTranslationEffect CSSName="left" DisplayName="X Position Animation" MaxValue="5000" MinValue="-5000" From="-200" To="1400" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="textIn" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#topInfo" Duration="1000" Easing="easeOutElastic" Callback="null"><Target.Effects><XTranslationEffect CSSName="left" DisplayName="X Position Animation" MaxValue="5000" MinValue="-5000" From="0" To="0" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /><TimerEffect CSSName="download" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="500" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><TimerEffect CSSName="littlePlane" DisplayName="Timer Effect" MaxValue="10000" MinValue="1" From="0" To="500" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="clearTimeout(timer);&#xD;&#xA;     timer = setTimeout(eval({1}),{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target><Target Name="#bubble" Duration="1000" Easing="linear" Callback="null"><Target.Effects><OpacityEffect CSSName="opacity" DisplayName="Opacity Animation" MaxValue="1" MinValue="0" From="0" To="0" IsStartValue="False" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target><Target Name="#bottomInfo" Duration="1000" Easing="linear" Callback="null"><Target.Effects><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="dayToNight" EventType="load" Trigger="{x:Null}"><Animation.Targets><Target Name="#sky" Duration="1000" Easing="linear" Callback="null"><Target.Effects><ColorEffect CSSName="color" DisplayName="Color Animation" MaxValue="1" MinValue="0" From="#ccffff" To="&quot;darkblue&quot;" IsStartValue="True" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="True" JQueryPluginURI="effects.core.js" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="download" EventType="[none]" Trigger="{x:Null}"><Animation.Targets><Target Name="#download" Duration="1000" Easing="linear" Callback="null"><Target.Effects><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="planeHover" EventType="mouseover" Trigger="#plane"><Animation.Targets><Target Name="#plane" Duration="1000" Easing="linear" Callback="null"><Target.Effects><OpacityEffect CSSName="opacity" DisplayName="Opacity Animation" MaxValue="1" MinValue="0" From="0" To="0" IsStartValue="False" IsActive="True" IsAnimatable="True" IsExpression="False" FormatString="" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="howdy2" EventType="mouseover" Trigger="#littleMushroomGuy"><Animation.Targets><Target Name="#howdy_bottom" Duration="1000" Easing="linear" Callback="null"><Target.Effects><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="visible" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation><Animation Name="howdy2_out" EventType="mouseout" Trigger="#littleMushroomGuy"><Animation.Targets><Target Name="#howdy_bottom" Duration="1000" Easing="easeOutElastic" Callback="null"><Target.Effects><ModifyCSSEffect CSSName="visibility" DisplayName="Modify CSS Effect" MaxValue="5000" MinValue="-5000" From="0" To="collapse" IsStartValue="False" IsActive="True" IsAnimatable="False" IsExpression="False" FormatString="$({0}).css({1},{2});&#xD;&#xA;" RequiresJQueryPlugin="False" JQueryPluginURI="" /></Target.Effects></Target></Animation.Targets></Animation></AnimationCollection>
$(document).ready(function ($) {
	var timer;

	function planeFly(event)
	{
	     $("#plane").css("top","0");
	     $("#plane").css("left","0");
	    $("#plane").animate({"top":-180,"left":820},5000, "easeInElastic", null);
	}

	function planedropTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("planeDrop"),"5000");
	}

	function planeDrop(event)
	{
	    $("#plane").animate({"top":0},300, "easeOutBounce", null);
	}

	function stripedKittyIn(event)
	{
		clearTimeout(timer);
	     timer = setTimeout(eval("pigIn"),"200");
	    $("#stripedKitty").css("left","-20").css("visibility","visible");
	    $("#stripedKitty").animate({"top":50,"left":550},800, "linear", null);
	}

	function stripedKittyInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("stripedKittyIn"),"400");
	 }

	function pigIn(event)
	{
		clearTimeout(timer);
	     timer = setTimeout(eval("whiteKittyIn"),"200");
	    $("#pig").css("right","-200").css("visibility","visible");
	    $("#pig").animate({"top":50,"right":1100},600, "linear", null);
	}

	function pigInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("pigIn"),"600");
	 }

	 function whiteKittyIn(event)
	{
	    clearTimeout(timer);
	    timer = setTimeout(eval("furryIn"),"200");
	    $("#whiteKitty").css("right","-50").css("visibility","visible");
	    $("#whiteKitty").animate({"top":10,"right":400},500, "linear", null);
	}

	function whiteKittyInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("whiteKittyIn"),"400");
	 }

	function furryIn(event)
	{
		clearTimeout(timer);
	     timer = setTimeout(eval("lizardIn"),"200");
	    $("#furry").css("right","-50").css("visibility","visible");
	    $("#furry").animate({"top":50,"right":900},400, "linear", null);
	}

	function furryInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("furryIn"),"400");
	 }

	function lizardIn(event)
	{
		clearTimeout(timer);
	     timer = setTimeout(eval("puppyIn"),"200");
	    $("#lizard").css("left","-50").css("visibility","visible");
	    $("#lizard").animate({"top":50,"left":950},600, "linear", null);
	}

	function lizardInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("lizardIn"),"600");
	 }

	function puppyIn(event)
	{
		clearTimeout(timer);
	     timer = setTimeout(eval("bunnyIn"),"200");
	    $("#puppy").css("left","-200").css("visibility","visible");
	    $("#puppy").animate({"top":50,"left":870},1400, "linear", null);
	}

	function puppyInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("puppyIn"),"700");
	 }

	 function bunnyIn(event)
	{
	    clearTimeout(timer);
	    timer = setTimeout(eval("girlPuppyIn"),"200");
	    $("#bunny").css("left","-50").css("visibility","visible");
	    $("#bunny").animate({"top":10,"left":350},800, "linear", null);
	}

	function bunnyInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("bunnyIn"),"400");
	 }

	 function girlPuppyIn(event)
	{
		
	    $("#girlPuppy").css("right","-200").css("visibility","visible");
	    $("#girlPuppy").animate({"top":100,"right":770},800, "linear", null);
	}

	function girlPuppyInTimer(event)
	{
	     clearTimeout(timer);
	     timer = setTimeout(eval("girlPuppyIn"),"300");
	 }


	planeFly();

	planedropTimer();

	planeDrop();

	girlPuppyInTimer();

	bunnyInTimer();

	puppyInTimer();

	lizardInTimer();

	furryInTimer();

	whiteKittyInTimer();

	pigInTimer();

	stripedKittyInTimer();

	dayToNight();


	$('#plane').bind('mouseover', planeHover);

	$('#littleMushroomGuy').bind('mouseover', howdy2);

	$('#littleMushroomGuy').bind('mouseout', howdy2_out);

});
