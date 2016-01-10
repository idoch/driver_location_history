'use strict';
var app = angular.module('driver_location_history');
app.controller('HomeController', ['$http',
  function ($http) {
    var self = this;
    self.searchedDriver = {};
    self.searchedEnv = 'IL';
    self.foundPoints = [];
    self.map_settings = {
      center: { latitude: 33.749990, longitude: -117.942114 },
      zoom: 13
    };

    self.openDatePicker = function (name) {
      self[name] = !self[name];
    };

    self.findDriverPoints = function () {
      if (self.searchedDriver == {}) return null;
      var env = self.searchedEnv;
      var driverId = self.searchedDriver.id;
      var params = {
        startDate: self.searchedDriver.startDate,
        endDate: self.searchedDriver.endDate
      };
      var url = '/api/v1/' + env + '/driver/' + driverId;
      $http({
        url: url,
        method: 'GET',
        params: params
      }).then(function (res) {
        var markers = _.map(res.data, function(item) {
          return {
            id: [item.DriverId],
            latitude: item.Lat,
            longitude: item.Lon
          };
        });
        self.foundPoints = markers;
      });
    };
}]);
