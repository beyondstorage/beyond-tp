import 'package:get/get.dart';
import 'package:ui/models/agents.dart';


class AgentsController extends GetxController {
  RxBool loading = false.obs;
  Rx<Agents> agents = Agents.fromList([]).obs;

  void getAgents() {
    // todo    Agent interface debugging
    loading(true);
    new Future.delayed(Duration(milliseconds: 1000)).then((value) {
      agents(Agents.fromList([
        {
          "name": "default",
          "id": "210003222",
          "ip": "192.168.1.1",
          "isOnline": true
        },
        {
          "name": "default",
          "id": "210002322",
          "ip": "192.168.2.1",
          "isOnline": false,
          "networkSpeed": 3.2,
          "taskNumber": 1
        }
      ]));
    });
  }
}
