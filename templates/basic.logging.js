Java.perform(() => {
    {{range $className,$class := .Profile.Classes -}}
    const {{escapeJS $className}} = Java.use('{{$className}}')
    {{range $constructor := $class.Constructors -}}
    {{escapeJS $className}}["$init"].overload({{getOverloadString $constructor.Arguments}}).implementation = function ({{getTypedArguments $constructor.Arguments}}) {
        console.log('{{$className}}.$init()')
        const returnValue = this["$init"].apply(this, arguments) // returns {{$className}}
        return returnValue
    }
    {{end}}
    {{range $methodName,$method := $class.Methods -}}
    {{range $overload := $method.Overloads -}}
    {{escapeJS $className}}["{{$methodName}}"].overload({{getOverloadString $overload.Arguments}}).implementation = function ({{getTypedArguments $overload.Arguments}}) {
        console.log('{{$className}}.{{$methodName}}()')
        const returnValue = this["{{$methodName}}"].apply(this, arguments) // returns {{$overload.ReturnType.ClassName}}
        return returnValue
    }
    {{end -}}
    {{end}}
    {{end}}
})