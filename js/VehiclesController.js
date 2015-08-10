trainingVehicles.controller('VehiclesCtrl', function VehiclesCtrl($scope, $http) {
  $http.get('http://127.0.0.1:8080/').then(function(resp) {
    console.log('Success', resp);
    // For JSON responses, resp.data contains the result
    $scope.vehicles = resp.data;
  }, function(err) {
    console.error('ERR', err);
    // err.status will contain the status code
  });
});
