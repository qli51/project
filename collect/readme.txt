ϵͳ���ؼ��ϵͳ

Agent�ˣ�

1���ռ�ϵͳָ��
��������Ϣ��������������ϵͳ��Ϣ
��������Ϣ��IP��ַ�б������д�ֽ�/������
��CPU��Ϣ��CPU�ͺš�CPU�߼�������CPU������ġ�CPUʹ���ʡ�
���ڴ�ʹ����������ڴ��С��ʵ���ڴ��С����������С
��Ӳ��ʹ���������Ӳ�̴�С����ʹ��Ӳ�̴�С
�޲ο��⣺ github.com/shirou/gopsutil

��ʱ�ϱ�
�ٽӿڼ�Ȩ�����̶�AppCode����Header�У�����Header����ӵ�Authorization�ֶΣ�����Authorization�ֶε�ֵΪ��APPCODE �� ��ǿո� ��APPCODEֵ�������� Authorization:APPCODE AppCodeֵ
��ÿ�����ϱ�һ��ϵͳָ��ṹ������

Server��
��1��֧��Agent�ϱ����ݽӿ�
�ٽӿڼ�Ȩ��У��http header�е�AppCodeֵ�Ƿ�һ��
���ϱ����ݳ־û��洢

��2��֧�ֶ��û���¼
���û���Ϣ��Ҫ���־û����û����벻�����Ĵ��
���û��˻�������ǰ�����ã�֧��admin��guest
�۳���¼�ӿ��⣬�����ӿ���Ҫ��Ȩ

��3���ṩ��ѯ�ϱ����ݽӿ�
�ٲ�ѯά�ȣ�������������ϵͳ���� Linux��Windows��Mac���ϱ�ʱ���


ʹ�÷�����

1.go run storeServer.go ���������
2.go run login.go ִ�е�¼����(��ִ�д˽ű���ִ�������ģ�����ʾ�ȵ�¼�������û���û�����Ϣ) �˻� admin ���� admin123
3.go run collectClient.go �����ɼ��붨ʱ�ϱ����ܣ��ṩ�־û��洢��dataRecordĿ¼��
4.go run getData.go -hostName xxx -osName xxx -startTime xxx -endTime xxx ���в���Ϊ��������

���Ҳ��ԣ��ɳɹ���ɿɳ������洢��������getData.go�ɹ���ȡ���˵���Ϣ


����Ŀ¼����
common �����Ŀ���õĹ�����
    - http.go ����http���ӵģ���ع�����װ
collect �����Ŀ��Ŀ¼
    - collectClient.go ��ʱ�ɼ����ϱ��ͻ���
    collectMethod �ɼ�����Ŀ¼
        - collectMethod.go �ɼ�����
        - collectMethod_test.go ����ű�
    common collect��Ŀ����ʹ��Ŀ¼
        - common.go ��Ҫ��¼�ӿڼ�Ȩ������¼�ķ���
    dataRecord �ɳ��������ݼ�¼·��
        - collectData ������ݼ�¼�ļ�
        - userInfo �˺������¼�ļ�
    - getData.go ���ⲿ���õģ���ѯ���ݵĽű�
    - login.go �ͻ��˵�¼�ű�
    - storeServer.go ���ڿ�������˼����Ľű�