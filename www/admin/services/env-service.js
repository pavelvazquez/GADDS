/**
 *
 */
window.disableThemeSettings = true;

angular.module('altrs.config.env', [])

// Register environment in AngularJS as constant
// '__env' should be defined in env.js
.constant('env', window.__env);
