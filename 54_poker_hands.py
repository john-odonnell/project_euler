# In Poker, a hand consists of 5 cards and ranked as such:
#
# Royal Flush: 10 J Q K A same suit                     9
# Straight Flush: 5 consecutive cards of same suit      8
# Four of a Kind: Four cards of the same value          7
# Full House: Three of a kind and a pair                6
# Flush: All cards of the same suit                     5
# Straight: All cards are consecutive values            4
# Three of a Kind: Three cards of the same value        3
# Two pairs: Two different pairs                        2
# One pair: Two cards of the same value                 1
# High card: Highest value card                         0
#
# The file poker.txt contains 1000 random hands dealt to two
# players. Formatted:
#
# 1a 1b 1c 1d 1e 2a 2b 2c 2d 2e
#
# How many hands does Player 1 win?

class Hand:
    """Represents a hand of poker"""

    card_value = {
        "2": 0, "3": 1, "4": 2, "5": 3, "6": 4,
        "7": 5, "8": 6, "9": 7, "T": 8, "J": 9,
        "Q": 10, "K": 11, "A": 12
    }

    card_suits = {
        "C": 0, "S": 1, "H": 2, "D": 3
    }

    def __init__(self, cards: list):
        self.cards = cards
        # used to determine winner between hands with same value
        self.first_in_straight = 0
        self.full_house = []
        self.four_of_a_kind = 0
        self.three_of_a_kind = 0
        self.pairs = []
        self.pair = 0
        # distributions of values and suits in hand
        #                      2  3  4  5  6  7  8  9 10  J  Q  K  A
        self.cards_by_value = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        #                     C  S  H  D
        self.cards_by_suit = [0, 0, 0, 0]
        # hand value from 0 (high card) to 9 (royal flush)
        self.hand_value = self.rank_hand()
        # hand values sorted from highest to lowest, used to break ties
        self.high_cards = self.highest_cards()

    def values_consecutive(self) -> tuple:
        """Determine if a hand is a straight"""
        consecutive = None
        first = 0
        count_cons = 0
        count_total = 0
        for val in self.cards_by_value:
            if consecutive == None and val == 1:
                consecutive = True
                count_cons += 1
                first = count_total
            elif consecutive == True and count_cons < 5:
                if val == 0:
                    consecutive = False
                else:
                    count_cons += 1
                count_total += 1
        return consecutive, first

    def highest_cards(self) -> list:
        """Sort the hand's cards from highest to lowest value"""
        sorted_cards = []
        for i in range(len(self.cards_by_value)-1, -1, -1):
            if self.cards_by_value[i] > 0:
                for j in range(0, self.cards_by_value[i]):
                    sorted_cards.append(i)
        return sorted_cards

    def rank_hand(self) -> int:
        """Rank the hand from 0 (high card) to 9 (royal flush)"""
        for card in self.cards:
            value = card[:1]
            suit = card[1:]
            self.cards_by_value[self.card_value[value]] += 1
            self.cards_by_suit[self.card_suits[suit]] += 1

        if 4 in self.cards_by_value:
            # four of a kind
            self.four_of_a_kind = self.cards_by_value.index(4) + 2
            return 7
        elif 3 in self.cards_by_value:
            if 2 in self.cards_by_value:
                # full house
                self.full_house.append(self.cards_by_value.index(3) + 2)
                self.full_house.append(self.cards_by_value.index(2) + 2)
                return 6
            else:
                # three of a kind
                self.three_of_a_kind = self.cards_by_value.index(3) + 2
                return 3
        elif 2 in self.cards_by_value:
            if self.cards_by_value.count(2) == 2:
                # two pair
                self.pairs = [i for i, x in enumerate(self.cards_by_value) if x == 2]
                return 2
            else:
                # pair
                self.pair = self.cards_by_value.index(2)
                return 1

        is_consecutive, first = self.values_consecutive()
        if is_consecutive:
            self.first_in_straight = first
            if 5 in self.cards_by_suit:
                if first + 2 == 10:
                    # royal flush
                    return 9
                else:
                    # straight flush
                    return 8
            else:
                # straight
                return 4
        else:
            if 5 in self.cards_by_suit:
                # flush
                return 5

        # high card
        return 0

file = open("poker.txt", "r")
hands = file.read().splitlines()
file.close()

player_1_wins = 0
player_2_wins = 0

def high_card_comp(player_1_hand: Hand, player_2_hand: Hand) -> bool:
    """returns True is player 1 beats player 2"""
    for i in range(0, 5):
        if player_1_hand.high_cards[i] > player_2_hand.high_cards[i]:
            return True
        elif player_2_hand.high_cards[i] > player_1_hand.high_cards[i]:
            return False
        else:
            i += 1

for hand in hands:
    hand = hand.split(" ")
    player_1_hand = Hand(hand[:5])
    player_2_hand = Hand(hand[5:])

    if player_1_hand.hand_value > player_2_hand.hand_value:
        player_1_wins += 1
    elif player_2_hand.hand_value > player_1_hand.hand_value:
        player_2_wins += 1
    else:
        if player_1_hand.hand_value == 0:
            if high_card_comp(player_1_hand, player_2_hand):
                player_1_wins += 1
            else:
                player_2_wins += 1
        if player_1_hand.hand_value == 1:
            if player_1_hand.pair > player_2_hand.pair:
                player_1_wins += 1
            elif player_2_hand.pair > player_1_hand.pair:
                player_2_wins += 1
            else:
                if high_card_comp(player_1_hand, player_2_hand):
                    player_1_wins += 1
                else:
                    player_2_wins += 1
        if player_1_hand.hand_value == 2:
            if player_1_hand.pairs[0] > player_2_hand.pairs[0]:
                player_1_wins += 1
            elif player_2_hand.pairs[0] > player_1_hand.pairs[0]:
                player_2_wins += 1
            else:
                if player_1_hand.pairs[1] > player_2_hand.pairs[1]:
                    player_1_wins += 1
                elif player_2_hand.pairs[1] > player_1_hand.pairs[1]:
                    player_2_wins += 1
                else:
                    if high_card_comp(player_1_hand, player_2_hand):
                        player_1_wins += 1
                    else:
                        player_2_wins += 1
        if player_1_hand.hand_value == 3:
            if player_1_hand.three_of_a_kind > player_2_hand.three_of_a_kind:
                player_1_wins += 1
            elif player_2_hand.three_of_a_kind > player_1_hand.three_of_a_kind:
                player_2_wins += 1
        if player_1_hand.hand_value == 4:
            if player_1_hand.first_in_straight > player_2_hand.first_in_straight:
                player_1_wins += 1
            elif player_2_hand.first_in_straight > player_1_hand.first_in_straight:
                player_2_wins += 1
        if player_1_hand.hand_value == 5:
            if high_card_comp(player_1_hand, player_2_hand):
                player_1_wins += 1
            else:
                player_2_wins += 1
        if player_1_hand.hand_value == 6:
            if player_1_hand.full_house[0] > player_2_hand.full_house[0]:
                player_1_wins += 1
            elif player_2_hand.full_house[0] > player_1_hand.full_house[0]:
                player_2_wins += 1
        if player_1_hand.hand_value == 7:
            if player_1_hand.four_of_a_kind > player_2_hand.four_of_a_kind:
                player_1_wins += 1
            elif player_2_hand.four_of_a_kind > player_1_hand.four_of_a_kind:
                player_2_wins += 1
        if player_1_hand.hand_value == 8:
            if player_1_hand.first_in_straight > player_2_hand.first_in_straight:
                player_1_wins += 1
            elif player_2_hand.first_in_straight > player_1_hand.first_in_straight:
                player_2_wins += 1

print(player_1_wins)

# CORRECT
