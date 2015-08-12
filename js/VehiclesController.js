trainingVehicles.config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    // route for the home page
    when('/vehicles/', {
        templateUrl : 'pages/vehicles.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/vehicles/:stocknumber/', {
        templateUrl : 'pages/details.html',
        controller  : 'VehiclesCtrl'
    });
}]);

trainingVehicles.controller('VehiclesCtrl', ['$scope', '$http', '$routeParams', function VehiclesCtrl($scope, $http, $routeParams) {
  $scope.stocknumber = $routeParams.stocknumber;
  
  $http.get('http://127.0.0.1:8080/api/vehicles/').then(function(resp) {
    console.log('Success', resp);
    // For JSON responses, resp.data contains the result
    $scope.vehicles = resp.data;
  }, function(err) {
    console.error('ERR', err);
    // err.status will contain the status code
  });
}]);
