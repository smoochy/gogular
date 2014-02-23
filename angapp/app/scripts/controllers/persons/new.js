'use strict';

angular.module('angappApp')
  .controller('NewPersonsCtrl', function ($scope, $routeParams, personsService) {
    var savePersonSuccess = function(response){
    	// what should I do here?
    };

    var savePersonError = function(response){
    	alert(response);
    };

    $scope.savePerson = function(person){
    	console.log(person);
    	var params = {
    		'first_name' : person.FirstName,
    		'last_name' : person.LastName
    	};
    	personsService.savePerson(params, savePersonSuccess, savePersonError);
    };

  });
