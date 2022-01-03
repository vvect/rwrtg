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
    {{range $className,$class := .Classes -}}
    const {{escapeJS $className}} = Java.use('{{$className}}')
    {{range $methodName,$method := $class.Methods -}}
    {{range $i, $overload := $method.ArgumentTypes -}}
    {{escapeJS $className}}["{{$methodName}}"].overload({{getOverloadString $overload}}).implementation = function ({{getTypedArguments $overload}}) {
        console.log('{{$className}}.{{$methodName}}(' + parametersToStrings([{{getTypedArguments $overload}}]) + ')')
        {{if hasReturnValue $methodName $method.ReturnTypes -}}
        const returnValue = this["{{$methodName}}"].apply(this, arguments)  // returns {{getReturnValueType $class $method $method.ReturnTypes $i}}
        return returnValue
    {{end -}}
    }
    {{end -}}
    {{end -}}
    {{end}}
})