var app = angular.module('app', ['ui.router']);

app.config(function($stateProvider, $urlRouterProvider) {

    $urlRouterProvider.otherwise('/home');
    $stateProvider
        .state('login', {
            url: '/login',
            templateUrl: 'index.html',
            controller: mainCtrl
        })

        .state('register', {
            url: '/register',
            templateUrl: 'register.html',
            controller: registerCtrl
        });
});

app.controller("mainCtrl",function ($scope,$http) {
    $scope.result = "";
    $scope.username="";
    $scope.password="";
    $scope.authenticated = false;
    $scope.login = function () {
        var jData =
            JSON.stringify({
                username: $scope.username,
                password:$scope.password
            });
        $http.post("http://localhost:8080/login", jData).success(function(res, status) {
            $scope.username = "";
            $scope.password = "";
            if(res == "true")
            {
                $scope.authenticated =true;
                $scope.result = "Welcome Back"
            }
            else
            {
                alert("Wrong Username or Password")
            }
        })
    }
});
app.controller("registerCtrl",function ($scope,$http) {
    $scope.model = {
        name:"",
        username:"",
        password:""
    }
    $scope.result = "";
    $scope.authenticated = false;
    $scope.login = function () {
        var jData =
            JSON.stringify({
                username: $scope.model.username,
                password:$scope.model.password,
                name:$scope.model.name
            });
        $http.post("http://localhost:8080/login", jData).success(function(res, status) {
            $scope.model.username = "";
            $scope.model.password = "";
            $scope.model.name = "";

            if(res == "true")
            {
                $state.go('login')
            }
            else
            {
                alert("Wrong Username or Password")
            }
        })

    }
});