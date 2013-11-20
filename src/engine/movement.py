from lib import contract, base

from . import file_loader, tile


class Movement(base.BaseDictionary):
    pass


@contract.accepts(tile.Tiles)
def load_movements(tiles):
    @contract.accepts(str)
    @contract.returns(Movement)
    def movement_getter(name):
        movement_ = args[name]
        speeds = {}
        for tile_, multiplier in movement_['speeds'].items():
            speeds[tiles[str(tile_)]] = multiplier
        return Movement(speeds)

    args = file_loader.load_struct('movements')
    return movement_getter, args.keys()
