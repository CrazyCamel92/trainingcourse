var app = angular.module('app', ['ui.router']);

app.config(function($stateProvider, $urlRouterProvider) {

    $urlRouterProvider.otherwise('/login');
    $stateProvider
        .state('login', {
            url: '/login',
            templateUrl: 'login.html'
        })

        .state('register', {
            url: '/register',
            templateUrl: 'register.html'
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
            if(res == "false")
            {
                alert("Wrong Username or Password")
            }
            else
            {
                $scope.authenticated =true;
                $scope.result = res;
            }
        })
    }
});
app.controller("registerCtrl",function ($scope,$http,$state) {
    $scope.model = {
        name:"",
        username:"",
        password:""
    }
    $scope.result = "";
    $scope.authenticated = false;
    $scope.register = function () {
        var jData =
            JSON.stringify({
                credentials:{ username: $scope.model.username,
                            password:$scope.model.password},
                name:$scope.model.name
            });
        $http.post("http://localhost:8080/register", jData).success(function(res, status) {
            $scope.model.username = "";
            $scope.model.password = "";
            $scope.model.name = "";


            if(res == "true")
            {
                $state.go('login')
            }
            else
            {
                alert(res)
            }
        })

    }
});