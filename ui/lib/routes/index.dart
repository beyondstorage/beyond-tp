import 'package:get/get.dart';

// import '../modules/endpoints/index.dart';
import '../modules/dashboard/index.dart';
import '../modules/service/index.dart';
import '../modules/signin/index.dart';
import '../modules/agents/index.dart';

class Routes {
  static final String main = "/";
  static final String login = "/login";
  static final String agents = "/agents";
  static final String services = "/services";

  static final List<GetPage> getPages = [
    GetPage(name: main, page: () => Dashboard()),
    GetPage(name: services, page: () => Service()),
    GetPage(name: agents, page: () => Agents()),
    GetPage(name: login, page: () => SignIn()),
  ];
}
