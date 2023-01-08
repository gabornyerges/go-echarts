package templates

var BaseTpl = `
{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>

<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};
	let action_{{ .ChartID | safeJS }} = {{ .JSONNotEscapedAction | safeJS }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
 	goecharts_{{ .ChartID | safeJS }}.dispatchAction(action_{{ .ChartID | safeJS }});

	var wsClient_{{ .ChartID | safeJS }} = new WebSocket("ws://localhost:8089/ws/{{ .ChartID| safeJS }}")
	wsClient_{{ .ChartID | safeJS }}.onopen = function() { console.log("websocket connection has opened."); }
	wsClient_{{ .ChartID | safeJS }}.onmessage = function(event) { goecharts_{{ .ChartID | safeJS }}.setOption(JSON.parse(event.data)) }
	wsClient_{{ .ChartID | safeJS }}.onclose = function() { console.log("websocket connection has been closed.") }

    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
{{ end }}
`
