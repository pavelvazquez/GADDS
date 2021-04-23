/**
 * @class LoginController
 * @ngInject
 */
function LoginController(UserService, $state) {
  var ctl = this;

  ctl.config = null;

  // ApiService.getConfig().then(function(config){
  //   ctl.config = config;
  // });

  ctl.signUp = function(user){
    return UserService.signUp(user)
    .then(function(){
      $state.go('app.info');
    });
  };

}


angular.module('altrs.controller.login', ['altrs.service.user'])
.controller('LoginController', LoginController);
