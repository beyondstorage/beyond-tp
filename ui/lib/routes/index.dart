import 'package:get/get.dart';

// import '../modules/endpoints/index.dart';
import '../modules/dashboard/index.dart';
import '../modules/signin/index.dart';

class Routes {
  static final String main = "/";
  static final String login = "/login";
  static final String endpoints = "/endpoints";

  static final List<GetPage> getPages = [
    GetPage(name: main, page: () => Dashboard()),
    // GetPage(name: endpoints, page: () => Endpoints()),
    GetPage(name: login, page: () => SignIn()),
  ];
}
