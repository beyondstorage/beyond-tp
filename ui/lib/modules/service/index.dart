import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../layout/index.dart';
import '../../widgets/page_description/index.dart';
import '../../widgets/empty_entry_list/index.dart';

import 'entry_list.dart';
import 'create_service_dialog/index.dart';
import 'controller.dart';

class Service extends GetView<ServiceController> {
  final ServiceController c = Get.put(ServiceController());

  Service() {
    c.getIdentities();
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      child: Column(
        children: [
          PageDescription(
            icon: IconData(0xe60b, fontFamily: 'tpIcon'),
            title: 'Services'.tr,
            subtitle:
                "Support binding one or more cloud service accounts / API key"
                    .tr,
          ),
          Obx(() => controller.identities.value.length() == 0
              ? EmptyEntryList(
                  icon: IconData(0xe60b, fontFamily: 'tpIcon'),
                  title: 'The service list is empty'.tr,
                  subTitle:
                      'Please click the button below to create service'.tr,
                  buttonText: 'Create service'.tr,
                  onClick: () => Get.dialog(CreateServiceDialog(
                      getIdentities: controller.getIdentities)),
                )
              : EntryList()),
        ],
      ),
    );
  }
}
