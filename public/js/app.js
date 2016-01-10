// Declare app level module which depends on filters, and services
'use strict';
angular.module('driver_location_history', ['ngResource',
                                           'ngRoute',
                                           'ui.bootstrap',
                                           'uiGmapgoogle-maps'])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/home/home.html',
        controller: 'HomeController',
        controllerAs: 'home'
      })
      .otherwise({redirectTo: '/'});
  }]);
