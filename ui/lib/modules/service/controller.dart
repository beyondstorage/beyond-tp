import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../models/service.dart';
import '../../common/request.dart';

class ServiceController extends GetxController {
  RxBool loading = false.obs;
  Rx<Services> identities = Services.fromList([]).obs;

  final String query =
      '''
    query {
      identities(type: Qingstor) {
        name
        type
        credential {
          protocol
          args
        }
        endpoint {
          protocol
          host
          port
        }
      }
    }
  ''';

  void getIdentities() {
    loading(true);

    queryGraphQL(QueryOptions(document: gql(query))).then((result) {
      loading(false);

      if (result.data != null) {
        identities(Services.fromList(result.data["identities"] ?? []));
      }
    }).catchError((error) {
      loading(false);
    });
  }

  Future<QueryResult> deleteIdentity(Service service) {
    String _query =
        '''
      mutation {
        deleteIdentity(input: { name: "${service.name}", type: ${service.type} }) { }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      getIdentities();

      return result;
    });
  }
}
