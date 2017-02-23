kd = require 'kd'


fn = (err) ->
  return no  unless err

  if Array.isArray err
    @fn er  for er in err
    return err.length

  content = 'An error occured!'
  if 'string' is typeof err
    content = err
    err = {}
  else if 'object' is typeof err
    content = err.message  if err.message


  err.type                   ?= 'default'
  err.content                ?= content
  err.dismissible            ?= yes
  err.duration               ?= 3000
  err.primaryButtonTitle     ?= null
  err.onPrimaryButtonClick   ?= null
  err.secondaryButtonTitle   ?= null
  err.onSecondaryButtonClick ?= null

  { addNotification } = kd.singletons.notificationViewController
  addNotification err

  yes


module.exports = fn
