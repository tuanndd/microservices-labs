import sys
import time

from oslo_log import log as logging
from oslo_config import cfg
import oslo_messaging as messaging

LOG = logging.getLogger(__name__)

quiet = False

_options = [
    cfg.StrOpt('name',
               help='Name of this server (used for direct addressing)',
               required=True),
    cfg.StrOpt('topic',
               default='my-topic',
               help="target topic, default 'my-topic'"),
    cfg.StrOpt('exchange',
               default="my-exchange",
               help="target exchange, default 'my-exchange'"),
    cfg.StrOpt('namespace',
               help="target namespace, default None"),
    cfg.StrOpt("url",
               required=True,
               default="rabbit://localhost:5672",
               help="transport address, default 'rabbit://localhost:5672'"),
    cfg.StrOpt("target_version",
               help="Override the default version in the target address"),
    cfg.BoolOpt("quiet",
                default=False,
                help="Suppress all stdout output"),
    cfg.StrOpt("log_levels",
               help="Set module specific log levels, e.g."
               " 'amqp=WARN,oslo.messaging=INFO,...'")
]


class TestEndpoint(object):
    global quiet
    def __init__(self, server, target=None):
        self.server = server
        self.target = target

    def test_method1(self, ctx, **args):
        res = "Result from test_method1" + str(args)
        print(res)
        return res
    
    def test_method2(self, ctx, **args):
        res = "Result from test_method2" + str(args)
        print(res)
        return res


def main(argv=None):

    global quiet

    logging.register_options(cfg.CONF)
    cfg.CONF.register_cli_opts(_options)
    cfg.CONF(sys.argv[1:])
    if cfg.CONF.log_levels:
        logging.set_defaults(
            default_log_levels=cfg.CONF.log_levels.split(',')
        )
    logging.setup(cfg.CONF, "server")

    quiet = cfg.CONF.quiet
    server_name = cfg.CONF.name
    exchange = cfg.CONF.exchange
    topic = cfg.CONF.topic
    namespace = cfg.CONF.namespace
    url = cfg.CONF.url
    target_version = cfg.CONF.target_version

    transport = messaging.get_rpc_transport(cfg.CONF, url=url)

    target = messaging.Target(exchange=exchange,
                              topic=topic,
                              namespace=namespace,
                              server=server_name,
                              version=target_version)
    server = messaging.get_rpc_server(transport, target,
                                      [TestEndpoint(server_name, target)])

    server.start()
    if not quiet:
        print("Running server, name=%s exchange=%s topic=%s namespace=%s"
              % (server_name, exchange, topic, namespace))
    
    server.wait()
    

if __name__ == "__main__":
    sys.exit(main())