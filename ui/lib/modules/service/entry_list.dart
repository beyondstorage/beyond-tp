import 'package:get/get.dart';
import 'package:flutter/material.dart';

import 'controller.dart';
import 'toolbar.dart';
import 'panel.dart';

class EntryList extends GetView<ServiceController> {
  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Obx(
        () => Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SizedBox(height: 8),
            Toolbar(),
            SizedBox(height: 16),
            Expanded(
              child: LimitedBox(
                maxWidth: double.infinity,
                maxHeight: double.infinity,
                child: Scrollbar(
                  child: SingleChildScrollView(
                    child: Wrap(
                      spacing: 24,
                      runSpacing: 24,
                      children: [
                        ...controller.identities.value.services.map(
                          (service) => ServicePanel(
                            service: service,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
