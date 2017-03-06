var app = angular.module("app",[]);
app.controller("mainCtrl",function ($scope,$http) {
    $scope.result = "";
    $scope.username="";
    $scope.password="";

    $scope.login = function () {
        var jData =
            JSON.stringify({
                username: $scope.username,
                password:$scope.password
            });
        console.log(jData)
        $http.post("http://localhost:8080/login", jData).success(function(jData, status) {
            $scope.result = data;
        })

    }
});