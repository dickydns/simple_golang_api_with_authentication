Feature: User Api

  Scenario: Mengambil data user
    When ambil data user
    Then status harus 200
    And response memiliki user list

  Scenario: Register data user
    Given payload data user
    When membuat data user
    Then status create harus 201
    And response data user