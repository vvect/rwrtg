function hasToString(handle){
    if (handle == null || handle == undefined){
        return "null"
    }

    if (handle.toString != null && handle.toString != undefined){
        return handle.toString()
    }else{
        return handle.$className
    }
}

function parametersToStrings(parameters){
    const stringParameters = parameters.map(
        param => {
            return hasToString(param)
        }
    )
    return stringParameters.join(', ')
}

console.log('')
console.log('The advanced logging template may cause more frequent crashes than the basic one, if you need to')
console.log('hook more classes, use the basic logging template and prune whatever classes you don\'t need.')

Java.perform(() => {
    {{range $className,$class := .Profile.Classes -}}
    const {{escapeJS $className}} = Java.use('{{$className}}')
    {{range $constructor := $class.Constructors -}}
    {{escapeJS $className}}["$init"].overload({{getOverloadString $constructor.Arguments}}).implementation = function ({{getTypedArguments $constructor.Arguments}}) {
         console.log('{{$className}}.$init(' + parametersToStrings([{{getTypedArguments $constructor.Arguments}}]) + ')')
        const returnValue = this["$init"].apply(this, arguments) // returns {{$className}}
        return returnValue
    }
    {{end}}
    {{range $methodName,$method := $class.Methods -}}
    {{range $overload := $method.Overloads -}}
    {{escapeJS $className}}["{{$methodName}}"].overload({{getOverloadString $overload.Arguments}}).implementation = function ({{getTypedArguments $overload.Arguments}}) {
        console.log('{{$className}}.{{$methodName}}(' + parametersToStrings([{{getTypedArguments $overload.Arguments}}]) + ')')
        const returnValue = this["{{$methodName}}"].apply(this, arguments) // returns {{$overload.ReturnType.ClassName}}
        return returnValue
    }
    {{end -}}
    {{end}}
    {{end}}
})