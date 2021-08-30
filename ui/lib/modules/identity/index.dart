import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../layout/index.dart';
import '../../widgets/page_description/index.dart';
import '../../widgets/empty_entry_list/index.dart';

import 'entry_list.dart';
import 'create_identity_dialog/index.dart';
import 'controller.dart';

class Identity extends GetView<IdentityController> {
  final IdentityController c = Get.put(IdentityController());

  Identity() {
    c.getIdentities();
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      child: Column(
        children: [
          PageDescription(
            icon: IconData(0xe60b, fontFamily: 'tpIcon'),
            title: 'Identities'.tr,
            subtitle:
                "Support Binding One Or More Cloud Service Accounts / API Key"
                    .tr,
          ),
          Obx(() => controller.identities.value.length() == 0
              ? EmptyEntryList(
                  icon: IconData(0xe60b, fontFamily: 'tpIcon'),
                  title: 'The Identity List Is Empty'.tr,
                  subTitle:
                      'Please Click The Button Below To Create Identity'.tr,
                  buttonText: 'Create Identity'.tr,
                  onClick: () => Get.dialog(CreateIdentityDialog(
                      getIdentities: controller.getIdentities)),
                )
              : EntryList()),
        ],
      ),
    );
  }
}
