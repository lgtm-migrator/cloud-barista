# v.0.6.0-CafeLatte (2022.07.08)
### API Change
- cb-tumblebug 변경된 API 반영

### Feature
- MCIS/VM 생성 시 최적배치 기능 추가
- MCIS 동적 생성 기능 추가
- 리소스 Inspecting 기능 추가

### Bug Fix
- CSP 아이콘 출력 방식 개선
- MCKS Add/Delete Node 보완



# v.0.5.0-Affogato (2021.12.16)
### API Change
- 성능개선 : MCIS/MCKS Manage에 go routin 적용
- cb-spider v0.4.19, cb-tumblebug v0.4.16, cb-mcks v0.4.6, cb-dragonfly v0.4.4 버전의 변경된 API 반영

### Feature
- 작업내역 표시 추가
- MCIS, VM Create 고도화
- 처리결과 return 방식 통합
- 지도 표시 추가

### Bug Fix
- chart 표시 보완



# v0.4.0-CafeMocha (2021.06.30)
### API Change
- 모니터링 Ploicy, Threshold 추가
- UI에서 직접 Framework 직접호출 방식 -> go server를 통해 호출하는 방식으로 변경

### Feature
- MCIS, VM 생성 시 Import/Export 기능 추가
- VM 생성 시 Expert 기능 추가
- cb-tumblebug & cb-spider & cb-dragonfly & cb-ladybug 변경된 API 반영
- credential에 provider별 설정기능 보완
- table 검색, 정렬기능 추가
- Main화면 추가

### Bug Fix
- style 보완 및 validation 고도화



# v0.3.0-espresso (2020.12.10.)
### API Change
- cb-tumblebug v0.2.9 / cb-dragonfly v0.2.8 버전의 변경된 API 반영
  - 호환성 테스트 완료 (각 프레임워크의 v0.3.0-espresso 버전과 동일)
### Feature
- 신규 디자인 반영및 구조 변경
- MCIR (Network / Security / SSH Key / Image / Spec) 기능 추가
- MCIS 생성 시 모니터링 에이전트 자동 설치

### Bug Fix
- 최초 로그인 시 NS 목록 조회 오류 메시지 출력



# v0.2.0-cappuccino (2020.06.03.)

### API Change
- Geolocation API 추가
- 전체 Request 및 Response Body의 상세 항목 변경
- 각 MCIS에서 PublicIP 추출 기능 API 추가
- Common URL API 추가

### Feature
- Location에 선택된 서비스의 위치 반영
- VM 모니터링 활성화 & 모니터링 차트 추가
- Dashboard 변경및 메인 내용 영문화 적용
- 환경 & 리소스 설정 기능 변경 및 보완
- 환경 변수에 로그인 계정 설정 추가
- cb-tumblebug & cb-spider & cb-dragonfly 변경된 API 반영

### Bug Fix
- 환경 & 리소스 설정 버그 수정



# v0.1.0-americano (2019.12.23.)

### Feature
- 모니터링 일부 연동
- 개별 및 서비스 라이프 사이클 제어
- 멀티 클라우드 서비스 생성 기능
- 일부 환경 & 리소스 설정 기능
