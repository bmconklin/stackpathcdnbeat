from stackpathcdnbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Stackpathcdnbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        stackpathcdnbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("stackpathcdnbeat is running"))
        exit_code = stackpathcdnbeat_proc.kill_and_wait()
        assert exit_code == 0
