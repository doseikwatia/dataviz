requirejs.config({
  baseUrl: '/assets/js',
  shim: {
    bootstrap: {
      deps: ['jquery', 'popper']
    }
  },
  paths: {
    jquery: 'jquery-3.2.1.slim.min',
    knockout: 'knockout-3.4.2',
    text: 'text',
    bootstrap: 'bootstrap.min',
    popper: 'popper.min'
  }
});

requirejs(['jquery','knockout'],function ($,ko) {

  //Binding application view model to dom when document is ready
  $('document').ready(function(){

  })
})