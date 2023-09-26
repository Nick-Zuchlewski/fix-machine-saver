import logging
import time
from pymodbus.client import ModbusTcpClient
from pymodbus.constants import Endian
from pymodbus.payload import BinaryPayloadBuilder
from pymodbus.payload import BinaryPayloadDecoder

# Logger
logging.basicConfig(format='%(levelname)s %(asctime)s %(message)s')
logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

def read_holding_registers(client: ModbusTcpClient, slave_id: int, address: int, count: int) -> None:
    """
    Read holding registers
    """
    try:
        result = client.read_holding_registers(address, count, unit=slave_id)
        if result.isError():
            logger.error(result)
        else:
            logger.info(result.registers)
    except Exception as e:
        logger.error(e)

def main() -> None:
    slave_ids = [38, 68, 61, 57, 94]

    try:
        logger.info("connecting...")
        client = ModbusTcpClient(host='10.111.100.182', port=502)
    except Exception as e:
        logger.error(e)
        logger.info("exiting...")
        return

    logger.info("connected")

    while True:
        # Read holding registers of each slave ID    
        for slave_id in slave_ids:
            logger.info("reading holding registers of slave ID: %d", slave_id)
            read_holding_registers(client, slave_id, 176, 3)
            time.sleep(1.0)
        # Sleep till next cycle
        logger.info("sleeping...")
        time.sleep(1.0)
    

if __name__ == "__main__":
    main()