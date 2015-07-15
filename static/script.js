
var button = angular.module("button", []);

button.controller('mainCtrl', function($scope, $http) {
	$scope.clickedButton = function() {
		$http.get('/buttonClicked').
		  success(function(data, status, headers, config) {
		    $scope.clickCount = data.clicks;
		  }).
		  error(function(data, status, headers, config) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		  });
	};
})

