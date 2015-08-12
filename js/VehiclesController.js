trainingVehicles.config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    // route for the home page
    when('/list', {
        templateUrl : 'pages/list.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/2003F1', {
        templateUrl : 'pages//2003F1.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/CM989', {
        templateUrl : 'pages/CM989.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/CM999', {
        templateUrl : 'pages/CM999.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/LAMBO', {
        templateUrl : 'pages/LAMBO.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/PW351', {
        templateUrl : 'pages/PW351.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/PW7122', {
        templateUrl : 'pages/PW7122.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/PW7165', {
        templateUrl : 'pages/PW7165.html',
        controller  : 'VehiclesCtrl'
    }).

    when('/06SaleenS7', {
        templateUrl : 'pages/06SaleenS7.html',
        controller  : 'VehiclesCtrl'
    });
}]);

trainingVehicles.controller('VehiclesCtrl', function VehiclesCtrl($scope, $http) {
  $http.get('http://127.0.0.1:8080/api/vehicles/').then(function(resp) {
    console.log('Success', resp);
    // For JSON responses, resp.data contains the result
    $scope.vehicles = resp.data;
  }, function(err) {
    console.error('ERR', err);
    // err.status will contain the status code
  });
});
