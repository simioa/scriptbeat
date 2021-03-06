import os
import sys
import unittest

sys.path.append(os.path.join(os.path.dirname(__file__), '../../tests/system'))
import metricbeat


@metricbeat.parameterized_with_supported_versions
class Test(metricbeat.BaseTest):

    COMPOSE_SERVICES = ['envoyproxy']

    @unittest.skipUnless(metricbeat.INTEGRATION_TESTS, "integration test")
    def test_stats(self):
        """
        EnvoyProxy module outputs an event.
        """
        self.render_config_template(modules=[{
            "name": "envoyproxy",
            "metricsets": ["server"],
            "hosts": self.get_hosts(),
            "period": "5s",
        }])
        proc = self.start_beat()
        self.wait_until(lambda: self.output_lines() > 0, max_timeout=20)
        proc.check_kill_and_wait()
        self.assert_no_logged_warnings()

        output = self.read_output_json()
        self.assertTrue(len(output) >= 1)
        evt = output[0]
        print(evt)

        self.assert_fields_are_documented(evt)
