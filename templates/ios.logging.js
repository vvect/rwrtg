{{range $className,$methods := .Classes -}}
const h{{escapeJS $className}} = ObjC.classes["{{$className}}"]
{{range $method := $methods}}
Interceptor.attach(h{{escapeJS $className}}["{{getIOSMethod $method}}"].implementation, {
    onEnter: function(args){
        console.log("h{{$className}}::{{getIOSMethod $method}}")
    },
    onLeave: function(retval){

    }
})
{{end}}
{{end}}